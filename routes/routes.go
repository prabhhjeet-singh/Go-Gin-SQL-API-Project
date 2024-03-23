package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvent)
	router.GET("/events/:id", getEventByID)

	authenticated := router.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", postEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	router.POST("/signup", userSignup)
	router.POST("/login", userLogin)
}
