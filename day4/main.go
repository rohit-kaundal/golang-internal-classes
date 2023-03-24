package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type Person struct {
	Fullname string
	DOB      int
}

func main() {
	fmt.Println("This is day 4")

	// http server , handling requests and responses
	// Using core golang http package

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte("{\"status\":\"up\"}"))
		return
	})

	// http endpoint using external handler function
	http.HandleFunc("/test", callBackTestCore)

	//http.ListenAndServe(":8080", nil)

	// Using gorilla mux
	router := mux.NewRouter()
	router.HandleFunc("/{userid}/{toudid}", callBackTestGorilla)
	router.HandleFunc("/post-test", callBackTestGorilla).Methods("POST")
	//http.ListenAndServe(":8080", router)

	// Using Gin
	ginRouter := gin.Default()
	ginRouter.POST("/", handleGinData)
	ginRouter.GET("/", handleGinData)
	ginRouter.POST("/login", processLoginData)
	ginRouter.Run()

	// Using fiber
	// DB persistence using core package
	// GORM , to persist data , CRUD operations into database
}

func callBackTestCore(rw http.ResponseWriter, r *http.Request) {
	log.Println(r)
	rw.Write([]byte("Hi welcome to callBackTest"))
}

func callBackTestGorilla(rw http.ResponseWriter, r *http.Request) {
	//log.Println(r)
	log.Println(mux.Vars(r))
	rw.Write([]byte("Request server through gorilla mux"))
}

func handleGinData(c *gin.Context) {
	type response struct {
		Firstname string `json:"first_name"`
		Dob       int    `json:"dob"`
	}
	//ginreposne := response{
	//	Firstname: "Deepak",
	//	Dob:       1976,
	//}
	//c.JSON(http.StatusOK, ginreposne)
	c.JSON(http.StatusInternalServerError, gin.H{"msg": "server-down"})
}

func processLoginData(c *gin.Context) {
	type LoginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	reqLogin := LoginData{}
	err := c.BindJSON(&reqLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid credentials"})
		return
	}
	log.Println("Username is ", reqLogin.Username, "with password", reqLogin.Password)
	c.JSON(http.StatusOK, gin.H{"msg": "successful login"})

}
