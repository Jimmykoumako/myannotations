package routers

import (
	"github.com/gin-gonic/gin"
	"your_project/controllers"
)

// SetUserRoutes sets up routes for user-related operations
func SetUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
		// Add other user routes as needed
	}
}
