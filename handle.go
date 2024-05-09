package main

import (
	"fmt"
	"net/http"

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
