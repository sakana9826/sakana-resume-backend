package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sakana9826/sakana-resume-backend/config"
	"github.com/sakana9826/sakana-resume-backend/models"
	"github.com/sakana9826/sakana-resume-backend/utils"
)

type GenerateAccessCodeRequest struct {
	ExpireHours int `json:"expireHours" binding:"required,min=1,max=168"`
}

type VerifyAccessCodeRequest struct {
	AccessCode string `json:"accessCode" binding:"required"`
}

func GenerateAccessCode(c *gin.Context) {
	var req GenerateAccessCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, err := utils.GenerateAccessCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access code"})
		return
	}

	expiresAt := time.Now().Add(time.Duration(req.ExpireHours) * time.Hour)
	accessCode := models.AccessCode{
		Code:      code,
		ExpiresAt: expiresAt,
	}

	if err := config.DB.Create(&accessCode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save access code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessCode": code,
		"expiresAt":  expiresAt,
	})
}

func VerifyAccessCode(c *gin.Context) {
	var req VerifyAccessCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var accessCode models.AccessCode
	if err := config.DB.Where("code = ?", req.AccessCode).First(&accessCode).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid access code"})
		return
	}

	if accessCode.Used {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access code has already been used"})
		return
	}

	if time.Now().After(accessCode.ExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access code has expired"})
		return
	}

	accessLog := models.AccessLog{
		AccessCode: accessCode.Code,
		IP:         c.ClientIP(),
		UserAgent:  c.Request.UserAgent(),
		AccessedAt: time.Now(),
		ExpiresAt:  accessCode.ExpiresAt,
	}

	if err := config.DB.Create(&accessLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access log"})
		return
	}

	accessCode.Used = true
	if err := config.DB.Save(&accessCode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update access code"})
		return
	}

	token, err := utils.GenerateToken(accessLog.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":     token,
		"expiresAt": accessCode.ExpiresAt,
	})
}
