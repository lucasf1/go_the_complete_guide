package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)    // GET, POST< PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // /events/1, /event/2

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	
	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	
	// server.POST("/events", middlewares.Authenticate, createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)


}
