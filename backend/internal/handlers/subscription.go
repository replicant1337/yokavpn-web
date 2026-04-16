package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"yokavpn-web-backend/internal/database"
	"yokavpn-web-backend/internal/models"
	"yokavpn-web-backend/internal/remnawave"

	"github.com/gin-gonic/gin"
)

type CreateSubRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func generateAuthKey() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func CreateSubscription(c *gin.Context) {
	var req CreateSubRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 0. Check or Create Local User
	var user models.User
	result := database.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		user = models.User{
			Email:    req.Email,
			Password: "system-generated", // Should be handled better in reality
		}
		if err := database.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create local user"})
			return
		}
	}

	client := remnawave.NewClient()

	// 1. Create User in Remnawave
	remnaUser, err := client.CreateUser(req.Username, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user in Remnawave: " + err.Error()})
		return
	}

	// 2. Create Subscription for that user
	remnaSub, err := client.CreateSubscription(remnaUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription in Remnawave: " + err.Error()})
		return
	}

	// 3. Save to local DB
	authKey := generateAuthKey()
	
	// Parse expiry date if needed
	expiresAt, _ := time.Parse(time.RFC3339, remnaSub.ExpiresAt)
	if remnaSub.ExpiresAt == "" {
		expiresAt = time.Now().AddDate(0, 1, 0) // Default 1 month
	}

	sub := models.Subscription{
		UserID:       user.ID,
		RemnaUserID:  remnaUser.ID,
		RemnaSubLink: remnaSub.SubLink,
		ShortID:      remnaSub.ShortID,
		AuthKey:      authKey,
		ExpiresAt:    expiresAt,
		Status:       "active",
	}

	if err := database.DB.Create(&sub).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save subscription locally"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":           "Subscription created successfully",
		"auth_key":          authKey,
		"short_id":          sub.ShortID,
		"subscription_link": sub.RemnaSubLink,
	})
}

func GetSubscriptionByAuthKey(c *gin.Context) {
	key := c.Param("authKey")
	if key == "" {
		key = c.Param("shortId")
	}
	
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key is required"})
		return
	}

	var sub models.Subscription
	if err := database.DB.Where("auth_key = ? OR short_id = ?", key, key).First(&sub).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	// Optionally sync with Remnawave here
	client := remnawave.NewClient()
	remnaSub, err := client.GetSubscriptionByShortID(sub.ShortID)
	if err == nil {
		// Update local cache if needed
		sub.RemnaSubLink = remnaSub.SubLink
		database.DB.Save(&sub)
	}

	c.JSON(http.StatusOK, sub)
}
