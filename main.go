package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID       int64  `json: "id"`
	UUID     string `json: "uuid"`
	NAME     string `json: "name"`
	PASSWORD string `json: "password"`
	PHONE    string `json: "phone"`
	STATUS   bool   `json: "status"`
}

var usersList = []Users{
	{ID: 1, UUID: "d24d4e54-3d3e-11ed-be6e-67909d485b89", NAME: "Janak", PASSWORD: "Test@1234", PHONE: "9865054974", STATUS: true},
	{ID: 2, UUID: "d685dfcc-3d3e-11ed-82b9-1f4e5fb4388d", NAME: "Jhon", PASSWORD: "Test@1234", PHONE: "9865054974", STATUS: true},
	{ID: 3, UUID: "da2c6204-3d3e-11ed-97ee-0b4d252c078b", NAME: "Ronit", PASSWORD: "Test@1234", PHONE: "9865054974", STATUS: true},
	{ID: 4, UUID: "dffb2080-3d3e-11ed-91a5-ab33c3e34fb2", NAME: "Racky", PASSWORD: "Test@1234", PHONE: "9865054974", STATUS: true},
}

func getUsersList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":    "operation successfull",
		"statusCode": 200,
		"data":       usersList,
	})
}

func createUser(c *gin.Context) {
	var input_users_data Users

	if err := c.ShouldBindJSON(&input_users_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	usersList = append(usersList, input_users_data)
	c.JSON(http.StatusOK, gin.H{
		"message":    "operation successfull",
		"statusCode": 200,
		"data":       usersList,
	})

}

func deleteUser(c *gin.Context) {
	uuid := c.Param("uuid")
	println(uuid)

	for index, element := range usersList {
		println(index)
		if uuid == element.UUID {
			c.JSON(http.StatusOK, gin.H{
				"message":    "delete successfull",
				"statusCode": 200,
			})

		}
		c.JSON(http.StatusNoContent, gin.H{
			"message": "operation failed",
		})

	}

}

func querystringParameters(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{

		v1.GET("/users", getUsersList)
		v1.POST("/create", createUser)
		v1.DELETE("/delete/:uuid", deleteUser)
		v1.GET("/info", querystringParameters)

	}

	println("Server is running on port:", 8080)
	http.ListenAndServe(":8080", router)
}
