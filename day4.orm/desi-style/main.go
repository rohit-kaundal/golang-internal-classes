package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm - Object relation mapping in golang (mysql, redis, postgres)

// Database
const DB_USERNAME = "root"
const DB_PASSWORD = "my_secret_password"
const DB_NAME = "gorm_tutorial"
const DNS = DB_USERNAME + ":" + DB_PASSWORD + "@tcp(127.0.0.1:6033)/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"

// structure which holds
type Task struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Username    string
	Title       string
	Description string
}

func main() {
	var db *gorm.DB
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&Task{})

	taskList := &[]Task{{
		ID:          2,
		Username:    "Deepak",
		Title:       "Learn golang",
		Description: "learning golang in next week",
	},
		{
			ID:          8,
			Username:    "Rohit Kapoor",
			Title:       "Learn gorm",
			Description: "Learn gorm",
		},
	}

	// Only Insert query
	//db.Create(&taskList)

	// Upsert Query - Insert when primary key is not given and update if primary key is given
	db.Save(&taskList)

	// Select * from Tasks
	var results *[]Task

	db.Find(&results)
	//fmt.Println(results)

	// Select * from Tasks where ID = 4
	var resultOfId *Task
	db.Where("id = ?", 6).First(&resultOfId)
	fmt.Println(resultOfId)

	// Deletion := for deletion initate the object and pass to delete function

	recordToDel := &Task{ID: 6}
	db.Delete(&recordToDel)

	// Our REST Api server using gin

	router := gin.Default()

	// Routes define

	// Get all tasks
	router.GET("/tasks", func(c *gin.Context) {
		var tasks *[]Task
		db.Find(&tasks)
		c.JSON(http.StatusOK, tasks)
	})

	// Get task by id
	router.GET("/task/:id", func(ctx *gin.Context) {
		var singleTask *Task
		taskId := ctx.Param("id")
		db.Where("id=?", taskId).First(&singleTask)
		ctx.JSON(http.StatusOK, singleTask)
	})

	// Create task
	router.POST("/task", func(c *gin.Context) {

		type userreq struct {
			Username  string `json:"user_name"`
			Tasktitle string `json:"title"`
			Taskdesc  string `json:"desc"`
		}

		var userReq userreq
		err := c.BindJSON(&userReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid request"})
			return
		}

		taskCreate := &Task{
			Username:    userReq.Username,
			Title:       userReq.Tasktitle,
			Description: userReq.Taskdesc,
		}
		db.Save(&taskCreate)
		c.JSON(http.StatusOK, gin.H{"msg": "Record added successfully"})

	})

	// Update task
	router.PUT("/task", func(c *gin.Context) {

		type userreq struct {
			ID        uint   `json:"user_id"`
			Username  string `json:"user_name"`
			Tasktitle string `json:"title"`
			Taskdesc  string `json:"desc"`
		}

		var userReq userreq
		err := c.BindJSON(&userReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid request"})
			return
		}

		taskCreate := Task{}
		taskCreate.ID = userReq.ID
		taskCreate.Username = userReq.Username
		taskCreate.Description = userReq.Taskdesc
		taskCreate.Title = userReq.Tasktitle

		log.Println(taskCreate)
		//Upsert query
		db.Where(Task{ID: taskCreate.ID}).Assign(taskCreate).FirstOrCreate(&taskCreate)
		c.JSON(http.StatusOK, taskCreate)

	})

	// Delete record by id
	router.DELETE("/task/:id", func(ctx *gin.Context) {
		taskID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

		db.Where(Task{ID: uint(taskID)}).Delete(&Task{})
		ctx.JSON(http.StatusOK, gin.H{"msg": "Record deleted"})
	})

	// Start server
	router.Run()
}
