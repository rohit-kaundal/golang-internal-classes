package service

import (
	"ddd_demo/models"
	"ddd_demo/repository"
	"errors"
	"fmt"
)

type ITaskManagerService interface {
	CreateTask(models.UserModel, models.TaskModel) error
}

type taskManagerService struct {
	repo repository.ITaskManagerRep
}

func NewTaskManagerService(repo repository.ITaskManagerRep) ITaskManagerService {
	return &taskManagerService{repo: repo}
}

func (t *taskManagerService) CreateTask(u models.UserModel, task models.TaskModel) error {
	// Logic to create the task
	fmt.Println("Created task:", task.TaskTitle, "for user:", u.FullName)
	if u.FullName == "" || task.TaskTitle == "" {
		return errors.New("invalid user or task")
	}
	if err := t.repo.CreateTaskInDB(u, task); err != nil {
		return err
	}
	return nil
}
