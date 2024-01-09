// router.go
package routers

import (
	"github.com/gin-gonic/gin"
	"mas/middleware"
	"mas/controllers"
	"mas/models"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	// Use the error handling middleware
	Router.Use(middleware.ErrorHandlerMiddleware())

	// Use the pagination middleware
	Router.Use(middleware.PaginationMiddleware())

	// Use the logging middleware
	Router.Use(middleware.LoggingMiddleware())

	// Use the input validation middleware
	// Router.Use(middleware.ValidationMiddleware(models.User{}))

	// Initialize controllers
	userController := controllers.UserController{}
	bookController := controllers.BookController{}
	textController := controllers.TextController{}
	annotationController := controllers.AnnotationController{}
	connectionController := controllers.ConnectionController{}

	// Public routes (no authentication required)
	Router.POST("/login", userController.LoginUser)
	Router.POST("/register", userController.RegisterUser)  // Added registration route


	// Apply auth middleware to routes that require authentication
	authGroup := Router.Group("/")
	// authGroup.Use(middleware.AuthMiddleware())

	// User-related routes
	authGroup.GET("/secure-endpoint", secureEndpointHandler)
	authGroup.GET("/user/:userID/books", middleware.PaginationMiddleware(), userController.ViewUserBooks)
	authGroup.POST("/user/:userID/books", userController.AddBookToUser)

	// Book-related routesk
	authGroup.POST("/books", middleware.ValidationMiddleware(models.Book{}), bookController.CreateBook)
	authGroup.GET("/books/:bookID", bookController.ViewBook)
	authGroup.PUT("/books/:bookID", bookController.UpdateBook)
	authGroup.DELETE("/books/:bookID", bookController.DeleteBook)
	authGroup.POST("/books/:bookID/texts", bookController.CreateTextInBook)

	// Text-related routes
	authGroup.POST("/texts", textController.CreateText)
	authGroup.GET("/texts/:textID", textController.ViewText)
	authGroup.PUT("/texts/:textID", textController.UpdateText)
	authGroup.DELETE("/texts/:textID", textController.DeleteText)
	authGroup.POST("/texts/:textID/annotations", textController.CreateAnnotationInText)

	// Annotation-related routes
	authGroup.POST("/annotations", annotationController.CreateAnnotation)
	authGroup.GET("/annotations/:annotationID", annotationController.ViewAnnotation)
	authGroup.PUT("/annotations/:annotationID", annotationController.UpdateAnnotation)
	authGroup.DELETE("/annotations/:annotationID", annotationController.DeleteAnnotation)
	authGroup.POST("/annotations/:annotationID1/connect/:annotationID2", annotationController.ConnectAnnotations)
	authGroup.GET("/annotations/:annotationID/connections", annotationController.ViewConnections)

	// Connection-related routes
	authGroup.POST("/connections", connectionController.CreateConnection)
	authGroup.GET("/connections/:connectionID", connectionController.ViewConnection)
	authGroup.PUT("/connections/:connectionID", connectionController.UpdateConnection)
	authGroup.DELETE("/connections/:connectionID", connectionController.DeleteConnection)
	authGroup.POST("/connections/:connectionID/feedback", connectionController.ProvideConnectionFeedback)
	authGroup.PUT("/connection-feedback/:feedbackID", connectionController.UpdateConnectionFeedback)
	authGroup.DELETE("/connection-feedback/:feedbackID", connectionController.DeleteConnectionFeedback)
	
}

func secureEndpointHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	c.JSON(200, gin.H{"message": "This is a secure endpoint", "user": user})
}
