package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvent)
	router.GET("/events/:id", getEventByID)
	router.POST("/events", postEvent)
	router.PUT("/events/:id", updateEvent)
	router.DELETE("/events/:id", deleteEvent)
	router.POST("/signup", userSignup)
	router.POST("/login", userLogin)
}
