package main

import (
	"net/http"

	"github.com/examples-hub/realworld-gin-gorm/middleware"
	"github.com/examples-hub/realworld-gin-gorm/models"
	"github.com/examples-hub/realworld-gin-gorm/router"
	"github.com/examples-hub/realworld-gin-gorm/validator"
	"github.com/gin-gonic/gin"
)

func main() {
	// config.InitConfig()
	models.InitDB()

	r := gin.Default()

	middleware.LoadMiddleware(r)
	validator.RegisteMyValidator(r)
	router.LoadRouter(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// addr := viper.GetString("serverAddr")
	// r.Run(addr)
	r.Run()
}
