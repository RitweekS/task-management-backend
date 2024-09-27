package routes

import (
	"task-management/controller"
	"task-management/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine){
	userController := controller.UserController{}
	taskController := controller.TaskController{}
	router.GET("/",userController.Check)

	v1:= router.Group("/v1")
	{
		v1.POST("/signup",userController.SignUp)
		v1.POST("/signin",userController.SignIn)

		task := v1.Group("/task")
		task.Use((middleware.JWTMiddleware()))
		{
			task.POST("/create",taskController.CreateTask)
			task.GET("/user/tasks",taskController.GetAllTaskByUser)
			task.GET("/:taskId", taskController.GetTaskById) 
			task.DELETE("/:taskId", taskController.DeleteTask) 
			task.PUT("/:taskId", taskController.UpdateTask) 
		}
	}
}
