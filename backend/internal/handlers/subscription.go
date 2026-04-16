package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
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
			Password: "system-generated", 
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
	remnaSub, err := client.CreateSubscription(remnaUser.ShortUuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription in Remnawave: " + err.Error()})
		return
	}

	// 3. Save to local DB
	var systemUser models.User
	if err := database.DB.Where("email = ?", "system@yokavpn.local").First(&systemUser).Error; err != nil {
		systemUser = models.User{
			Email:    "system@yokavpn.local",
			Password: "system-generated",
		}
		database.DB.Create(&systemUser)
	}

	authKey := generateAuthKey()
	expiresAt, _ := time.Parse(time.RFC3339, remnaSub.User.ExpiresAt)
	if remnaSub.User.ExpiresAt == "" {
		expiresAt = time.Now().AddDate(0, 1, 0)
	}

	trafficTotal, _ := strconv.ParseInt(remnaSub.User.TrafficLimitBytes, 10, 64)
	trafficUsed, _ := strconv.ParseInt(remnaSub.User.TrafficUsedBytes, 10, 64)

	newSub := models.Subscription{
		UserID:       user.ID,
		RemnaUserID:  remnaUser.ShortUuid,
		RemnaSubLink: remnaSub.SubscriptionUrl,
		ShortID:      remnaSub.User.ShortUuid,
		AuthKey:      authKey,
		TrafficTotal: trafficTotal,
		TrafficUsed:  trafficUsed,
		ExpiresAt:    expiresAt,
		Status:       "active",
	}

	if err := database.DB.Create(&newSub).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save subscription locally"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":           "Subscription created successfully",
		"auth_key":          authKey,
		"short_id":          newSub.ShortID,
		"subscription_link": newSub.RemnaSubLink,
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
	err := database.DB.Where("auth_key = ? OR short_id = ?", key, key).First(&sub).Error
	
	client := remnawave.NewClient()

	if err != nil {
		fmt.Printf("Subscription not found in DB for key %s, trying Remnawave...\n", key)
		remnaSub, err := client.GetSubscriptionByShortID(key)
		if err != nil {
			fmt.Printf("Remnawave lookup failed for key %s: %v\n", key, err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found locally or in Remnawave"})
			return
		}

		fmt.Printf("Found in Remnawave, importing: %+v\n", remnaSub.User.ShortUuid)
		
		var systemUser models.User
		if err := database.DB.Where("email = ?", "system@yokavpn.local").First(&systemUser).Error; err != nil {
			systemUser = models.User{
				Email:    "system@yokavpn.local",
				Password: "system-generated",
			}
			database.DB.Create(&systemUser)
		}

		authKey := generateAuthKey()
		expiresAt, _ := time.Parse(time.RFC3339, remnaSub.User.ExpiresAt)
		if remnaSub.User.ExpiresAt == "" {
			expiresAt = time.Now().AddDate(0, 1, 0)
		}

		trafficTotal, _ := strconv.ParseInt(remnaSub.User.TrafficLimitBytes, 10, 64)
		trafficUsed, _ := strconv.ParseInt(remnaSub.User.TrafficUsedBytes, 10, 64)

		sub = models.Subscription{
			UserID:       systemUser.ID,
			RemnaUserID:  remnaSub.User.ShortUuid,
			RemnaSubLink: remnaSub.SubscriptionUrl,
			ShortID:      remnaSub.User.ShortUuid,
			AuthKey:      authKey,
			TrafficTotal: trafficTotal,
			TrafficUsed:  trafficUsed,
			ExpiresAt:    expiresAt,
			Status:       "active",
		}

		if err := database.DB.Create(&sub).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to auto-import subscription"})
			return
		}
	} else {
		// Found locally, sync
		remnaSub, err := client.GetSubscriptionByShortID(sub.ShortID)
		if err == nil {
			trafficTotal, _ := strconv.ParseInt(remnaSub.User.TrafficLimitBytes, 10, 64)
			trafficUsed, _ := strconv.ParseInt(remnaSub.User.TrafficUsedBytes, 10, 64)
			
			sub.RemnaSubLink = remnaSub.SubscriptionUrl
			sub.TrafficTotal = trafficTotal
			sub.TrafficUsed = trafficUsed
			database.DB.Save(&sub)
		}
	}

	c.JSON(http.StatusOK, sub)
}
