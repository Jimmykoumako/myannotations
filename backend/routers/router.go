package routers

import (
   "github.com/gin-gonic/gin"
   "mas/controllers"
)

// InitializeRoutes sets up the application routes
func InitializeRoutes(userController *controllers.UserController, noteController *controllers.NoteController) *gin.Engine {
   router := gin.Default()

   // User routes
   userGroup := router.Group("/users")
   {
      userGroup.POST("/", userController.RegisterUser)
      // Add other user-related routes here
   }

   // Note routes
   noteGroup := router.Group("/notes")
   {
      noteGroup.POST("/", noteController.CreateNote)
      // Add other note-related routes here
   }

   return router
}
