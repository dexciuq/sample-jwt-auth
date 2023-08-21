package app

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(r *gin.RouterGroup) {
	api := r.Group("/api")
	api.GET("/healthcheck")

	// auth
	auth := api.Group("/auth")
	auth.POST("/register")
	auth.POST("/login")
	auth.POST("/verify")
	auth.GET("/refresh-token")
	auth.GET("/logout")

	// users
	users := api.Group("/users")
	users.GET("/me")

	// admin
	admin := api.Group("/admin")
	admin.GET("/users")
}
