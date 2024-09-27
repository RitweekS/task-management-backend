package controller

import (
	"net/http"
	"strconv"
	"task-management/model/common"
	"task-management/model/task"
	"task-management/service"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *service.TaskService
}

func (t *TaskController) CreateTask(cxt *gin.Context){
	userId, exists := cxt.Get("userId")
	if !exists {
       
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: "User ID not found",
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
        return
    }

	userIdInt, ok := userId.(int)
	if !ok {
		cxt.JSON(http.StatusInternalServerError, common.Response{
			ResponseMessage: "Failed to cast user ID",
			ResponseCode:    http.StatusInternalServerError,
			ResponseData:    nil,
		})
		return
	}


	var requestBody task.CreateTask
	requestBody.UserId = userIdInt

	if bindError := cxt.BindJSON(&requestBody);bindError!=nil{
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: bindError.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	data,err := t.taskService.CreateTask(requestBody)

	if err!=nil{
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: err.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	cxt.JSON(http.StatusOK,common.Response{
		ResponseMessage: "task created",
		ResponseCode: http.StatusOK,
		ResponseData: data,
	})

}
func (t *TaskController) GetAllTaskByUser(cxt *gin.Context){
	// var id = cxt.Param("id")
	
	// parseInt,parseIntError  := strconv.Atoi(id)

	
	// if parseIntError !=nil{
	// 	cxt.JSON(http.StatusBadRequest,common.Response{
	// 		ResponseMessage: parseIntError.Error(),
	// 		ResponseCode: http.StatusBadRequest,
	// 		ResponseData: nil,
	// 	})
	// 	return
	// }

	userId, exists := cxt.Get("userId")
	if !exists {
       
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: "User ID not found",
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
        return
    }

	parseInt, ok := userId.(int)
	if !ok {
		cxt.JSON(http.StatusInternalServerError, common.Response{
			ResponseMessage: "Failed to cast user ID",
			ResponseCode:    http.StatusInternalServerError,
			ResponseData:    nil,
		})
		return
	}
	


	data,err := t.taskService.GetAllTaskByUser(parseInt)

	if err!=nil{
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: err.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	cxt.JSON(http.StatusOK,common.Response{
		ResponseMessage: "",
		ResponseCode: http.StatusOK,
		ResponseData: data,
	})

}

func (t *TaskController) GetTaskById(cxt *gin.Context){
	var taskId = cxt.Param("taskId")
	// var userId = cxt.Param("userId")
	userId, exists := cxt.Get("userId")

	if !exists {
       
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: "User ID not found",
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
        return
    }

	userIdInt, ok := userId.(int)

	parseTaskId,parseTaskIdError  := strconv.Atoi(taskId)

	if !ok {
		cxt.JSON(http.StatusInternalServerError, common.Response{
			ResponseMessage: "Failed to cast user ID",
			ResponseCode:    http.StatusInternalServerError,
			ResponseData:    nil,
		})
		return
	}

	
	if parseTaskIdError != nil {
		errorMessage := ""
	
		if parseTaskIdError != nil {
			if errorMessage != "" {
				errorMessage += "; "
			}
			errorMessage += "Task ID Error: " + parseTaskIdError.Error()
		}
	
		cxt.JSON(http.StatusBadRequest, common.Response{
			ResponseMessage: errorMessage,
			ResponseCode:    http.StatusBadRequest,
			ResponseData:    nil,
		})
		return
	}

	data,err := t.taskService.GetTaskById(userIdInt,parseTaskId)
	if err!=nil{
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: err.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	cxt.JSON(http.StatusOK,common.Response{
		ResponseMessage: "",
		ResponseCode: http.StatusOK,
		ResponseData: data,
	})

}
func (t *TaskController) UpdateTask(cxt *gin.Context){
	var taskId = cxt.Param("taskId")
	// var userId = cxt.Param("userId")
	// parseUserId,parseUserIdError  := strconv.Atoi(userId)
	parseTaskId,parseTaskIdError  := strconv.Atoi(taskId)
	userId, exists := cxt.Get("userId")
	if !exists {
       
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: "User ID not found",
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
        return
    }


	parseInt, ok := userId.(int)
	if !ok {
		cxt.JSON(http.StatusInternalServerError, common.Response{
			ResponseMessage: "Failed to cast user ID",
			ResponseCode:    http.StatusInternalServerError,
			ResponseData:    nil,
		})
		return
	}
	
	
	if parseTaskIdError != nil {
		errorMessage := ""
		if parseTaskIdError != nil {
			if errorMessage != "" {
				errorMessage += "; "
			}
			errorMessage += "Task ID Error: " + parseTaskIdError.Error()
		}
	
		cxt.JSON(http.StatusBadRequest, common.Response{
			ResponseMessage: errorMessage,
			ResponseCode:    http.StatusBadRequest,
			ResponseData:    nil,
		})
		return
	}
	var requestBody task.UpdateTask
	if bindError := cxt.BindJSON(&requestBody);bindError!=nil{
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: bindError.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	data,err := t.taskService.UpdateTask(parseInt,parseTaskId,requestBody)
	if err!=nil{
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: err.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	cxt.JSON(http.StatusOK,common.Response{
		ResponseMessage: "Task updated",
		ResponseCode: http.StatusOK,
		ResponseData: data,
	})

}

func (t *TaskController) DeleteTask(cxt *gin.Context){
	var taskId = cxt.Param("taskId")
	// var userId = cxt.Param("userId")
	userId, exists := cxt.Get("userId")
	if !exists {
       
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: "User ID not found",
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
        return
    }

	parseInt, ok := userId.(int)
	// parseUserId,parseUserIdError  := strconv.Atoi(userId)
	parseTaskId,parseTaskIdError  := strconv.Atoi(taskId)

	if !ok {
		cxt.JSON(http.StatusInternalServerError, common.Response{
			ResponseMessage: "Failed to cast user ID",
			ResponseCode:    http.StatusInternalServerError,
			ResponseData:    nil,
		})
		return
	}
	

	
	if parseTaskIdError != nil {
		errorMessage := ""
		if parseTaskIdError != nil {
			if errorMessage != "" {
				errorMessage += "; "
			}
			errorMessage += "Task ID Error: " + parseTaskIdError.Error()
		}
	
		cxt.JSON(http.StatusBadRequest, common.Response{
			ResponseMessage: errorMessage,
			ResponseCode:    http.StatusBadRequest,
			ResponseData:    nil,
		})
		return
	}

	data,err := t.taskService.DeleteTask(parseInt,parseTaskId)
	if err!=nil{
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: err.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	cxt.JSON(http.StatusOK,common.Response{
		ResponseMessage: "Deleted Successfully",
		ResponseCode: http.StatusOK,
		ResponseData: data,
	})

}