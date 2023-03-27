package repository

import (
	"ddd_demo/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID        int
	TaskTitle string
	UserId    int
}

type ITaskManagerRep interface {
	CreateTaskInDB(models.UserModel, models.TaskModel) error
}

type taskManagerRepo struct {
	db *gorm.DB
}

func NewTaskMagaerSql(dsn string) ITaskManagerRep {
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	dbConnection.AutoMigrate(&Task{})

	return &taskManagerRepo{db: dbConnection}

}

func (svc *taskManagerRepo) CreateTaskInDB(u models.UserModel, t models.TaskModel) error {
	taskDb := Task{
		ID:        t.ID,
		TaskTitle: t.TaskTitle,
		UserId:    u.ID,
	}

	if err := svc.db.Where(Task{ID: taskDb.ID}).Assign(taskDb).FirstOrCreate(&taskDb).Error; err != nil {
		return err
	}
	return nil
}
