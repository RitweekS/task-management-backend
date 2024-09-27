package service

import (
	"fmt"
	"task-management/database"
	"task-management/model/task"
)

type TaskService struct {}

func (a *TaskService) CreateTask(tasks task.CreateTask)(bool,error){

	queryErr:= database.DB.Table("tasks").Create(&tasks)

	if(queryErr!=nil){
		fmt.Println("query error",queryErr)
		return false,queryErr.Error
	}

	return true,nil
}


func (a *TaskService) GetAllTaskByUser(id int)([]task.GetTask,error){

	var task = []task.GetTask{}
	queryErr:= database.DB.Table("tasks").Where("user_id = ?",id).Scan(&task)

	if(queryErr!=nil){
		fmt.Println("query error",queryErr)
		return task,queryErr.Error
	}

	return task,nil
}

func (a *TaskService) GetTaskById(userId int,taskId int)(task.GetTask,error){

	var task task.GetTask
	queryErr := database.DB.Table("tasks").Where("user_id = ? AND id = ?", userId, taskId).First(&task)

	if(queryErr!=nil){
		fmt.Println("query error",queryErr)
		return task,queryErr.Error
	}

	return task,nil
}

func (a *TaskService) UpdateTask(userId int,taskId int,tasks task.UpdateTask)(task.GetTask,error){

	var task task.GetTask
	queryErr := database.DB.Table("tasks").Where("user_id = ? AND id = ?", userId, taskId).Updates(tasks)

	if(queryErr.Error!=nil){
		fmt.Println("query error",queryErr)
		return task,queryErr.Error
	}

	getQueryErr := database.DB.Table("tasks").Where("user_id = ? AND id = ?", userId, taskId).First(&task)

	if(getQueryErr!=nil){
		fmt.Println("query error",queryErr)
		return task,queryErr.Error
	}
	

	return task,nil
}

func (a *TaskService) DeleteTask(userId int,taskId int)(bool,error){

	queryErr := database.DB.Table("tasks").Where("user_id = ? AND id = ?", userId, taskId).Delete(&task.CreateTask{})

	if(queryErr.Error!=nil){
		fmt.Println("query error",queryErr)
		return false,queryErr.Error
	}

	return true,nil
}