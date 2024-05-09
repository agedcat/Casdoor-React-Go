package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
)

func signinHandler(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	token, err := casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		fmt.Println("GetOAuthToken() error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get OAuth token"})
	}

	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)

	responsedata := map[string]interface{}{
		"status": "ok",
		"data":   token.AccessToken,
	}

	fmt.Println(responsedata)

	c.JSON(http.StatusOK, responsedata)
}

func getUserinfo(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authHeader is empty"})
	}

	token := strings.Split(authHeader, "Bearer ")
	if len(token) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid Bearer token"})
	}

	claims, err := casdoorsdk.ParseJwtToken(token[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse JWT token"})
	}

	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)

	responsedata := map[string]interface{}{
		"status": "ok",
		"data":   claims.User,
	}

	c.JSON(http.StatusOK, responsedata)
}
