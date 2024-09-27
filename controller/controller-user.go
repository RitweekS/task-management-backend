package controller

import (
	"net/http"
	"task-management/model/common"
	"task-management/model/user"
	"task-management/service"
	"task-management/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func (u *UserController) Check(cxt *gin.Context){
	

	cxt.JSON(http.StatusOK,common.Response{
		ResponseMessage: "working",
		ResponseCode: http.StatusOK,
		ResponseData: "",
	})
}
func (u *UserController) SignUp(cxt *gin.Context){
	var requestBody user.User

	if bindErr := cxt.BindJSON(&requestBody);bindErr!=nil{
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: bindErr.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	userCreated,err := u.userService.SignUp(requestBody)
	if err != nil {
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: err.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	cxt.JSON(http.StatusOK,common.Response{
		ResponseMessage: "user created",
		ResponseCode: http.StatusOK,
		ResponseData: userCreated,
	})
}

func (u *UserController) SignIn(cxt *gin.Context){
	var requestBody user.UserSignIn

	if bindErr := cxt.BindJSON(&requestBody);bindErr!=nil{
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: bindErr.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	userCreated,err := u.userService.SignIn(requestBody)
	if err != nil {
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: err.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
		return
	}

	token,tokenCreateErr := utils.CreateToken(userCreated.Id)
	if tokenCreateErr != nil {
		cxt.JSON(http.StatusBadRequest,common.Response{
			ResponseMessage: tokenCreateErr.Error(),
			ResponseCode: http.StatusBadRequest,
			ResponseData: nil,
		})
        return
    }

	cxt.JSON(http.StatusOK,common.Response{
		ResponseMessage: "successfully login!",
		ResponseCode: http.StatusOK,
		ResponseData: token,
	})
	

	// cxt.JSON(http.StatusOK,common.Response{
	// 	ResponseMessage: "unable to ",
	// 	ResponseCode: http.StatusOK,
	// 	ResponseData: userCreated,
	// })
	
}
