package routes

import (
	users_controller "argoCD-golang/src/user"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.POST("/", users_controller.InsertUser)
			users.GET("/", users_controller.FindUsers)
			users.PUT("/:id", users_controller.UpdateUser)
			users.DELETE("/:id", users_controller.DeleteUser)
		}

	}

	return router
}
