package service

import (
	"ddd_demo/models"
	"ddd_demo/repository"
	"testing"
)

var DSN string = "root:my_secret_password@tcp(127.0.0.1:6033)/ddd_demo?parseTime=True"

func TestCreateTask(t *testing.T) {
	user := models.UserModel{
		ID:       1,
		FullName: "Rohit Kaundal",
	}
	task := models.TaskModel{
		ID:        2,
		TaskTitle: "Dummy Task",
	}

	taskRepo := repository.NewTaskMagaerSql(DSN)
	taskManagerSvc := NewTaskManagerService(taskRepo)

	if err := taskManagerSvc.CreateTask(user, task); err != nil {
		t.Fail()
	}
}

func TestCreateInvalidTask(t *testing.T) {
	user := models.UserModel{
		ID:       1,
		FullName: "",
	}
	task := models.TaskModel{
		ID:        2,
		TaskTitle: "Dummy Task",
	}

	taskRepo := repository.NewTaskMagaerSql(DSN)
	taskManagerSvc := NewTaskManagerService(taskRepo)
	if err := taskManagerSvc.CreateTask(user, task); err == nil {
		t.Fail()
	}
}
