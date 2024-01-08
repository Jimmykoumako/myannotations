// routers/router.go

package routers

import (
    "github.com/gin-gonic/gin"
    "mas/controllers"
)

// SetupRouter initializes and returns the Gin router with defined routes.
func SetupRouter() *gin.Engine {
    router := gin.Default()

    api := router.Group("/api")
    {
        api.GET("/books", controllers.GetBooks)
        // Add more routes as needed
    }

    return router
}
