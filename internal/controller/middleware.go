package controller

import (
	"errors"
	"net/http"
	"strings"
	"warehouse/internal/service"
	"warehouse/utils"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userUsernameCtx     = "userUsername"
)

func getClaims(c *gin.Context) (*utils.CustomClaims, error) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return &utils.CustomClaims{}, errors.New("unauthorized")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return &utils.CustomClaims{}, errors.New("unauthorized")
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return &utils.CustomClaims{}, errors.New("unauthorized")
	}

	accessToken := headerParts[1]

	claims, err := utils.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return &utils.CustomClaims{}, errors.New("unauthorized")
	}

	return claims, nil
}

func CheckAdmin(c *gin.Context) {
	claims, err := getClaims(c)
	if err != nil {
		return
	}

	user, err := service.GetUserByID(claims.UserID)
	if err != nil {
		HandleError(c, err)
		return
	}

	if user.Role != "admin" || !user.Active {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "access denied",
		})
		return
	}

	// если что то понадобится то добавить в глобал через
	// фунцию set от context
	c.Set(userUsernameCtx, claims.Username)
	c.Next()
}

func CheckAuthorization(c *gin.Context) {
	claims, err := getClaims(c)
	if err != nil {
		return
	}

	user, err := service.GetUserByID(claims.UserID)
	if err != nil {
		HandleError(c, err)
		return
	}

	if !user.Active {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "access denied",
		})
		return
	}

	// если что то понадобится то добавить в глобал через
	// фунцию set от context
	c.Set(userUsernameCtx, claims.Username)
	c.Next()
}
