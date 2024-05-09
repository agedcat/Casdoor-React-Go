package main

import (
	"fmt"
	"log"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
)

func initAuthConfig() {
	casdoorsdk.InitConfig(
		GlobalConfig.Server.Endpoint,
		GlobalConfig.Server.ClientID,
		GlobalConfig.Server.ClientSecret,
		GlobalConfig.Certificate,
		GlobalConfig.Server.Organization,
		GlobalConfig.Server.Application,
	)
}

func main() {
	err := LoadConfig("app.yaml")
	if err != nil {
		panic(err)
	}

	initAuthConfig()

	r := gin.New()

	// 跨域设置
	r.Use(Cors())

	// 路由设置
	r.GET("/", func(c *gin.Context) {
		log.Println("home page")
	})

	r.POST("/api/signin", signinHandler)

	// 监听并在 0.0.0.0:8080 上启动服务
	fmt.Println("Server listening at: http://localhost:8080")
	r.Run(":8080")
}
