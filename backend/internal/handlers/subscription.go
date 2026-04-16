package handlers

import (
	"net/http"

	"yokavpn-web-backend/internal/remnawave"

	"github.com/gin-gonic/gin"
)

type CreateSubRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func CreateSubscription(c *gin.Context) {
	var req CreateSubRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := remnawave.NewClient()

	// 1. Create User in Remnawave
	remnaUser, err := client.CreateUser(req.Username, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user in Remnawave: " + err.Error()})
		return
	}

	// 2. Create Subscription for that user
	sub, err := client.CreateSubscription(remnaUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription in Remnawave: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":           "Subscription created successfully",
		"remna_user_id":     remnaUser.ID,
		"subscription_link": sub.SubLink,
		"short_id":          sub.ShortID,
	})
	}

	func GetSubscription(c *gin.Context) {
	shortID := c.Param("shortId")
	if shortID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short ID is required"})
		return
	}

	client := remnawave.NewClient()
	sub, err := client.GetSubscriptionByShortID(shortID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	c.JSON(http.StatusOK, sub)
	}
