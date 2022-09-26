package main

import (
	"GIN_REST_API/api/v1/routes"
	v1_routes "GIN_REST_API/api/v1/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/create", v1_routes.CreatePost)
	router.GET("read/:uuid", routes.ReadOnePost)
	router.PUT("/update/:uuid", routes.UpdatePost)
	router.DELETE("/delete/:uuid", routes.DeletePost)

	println("Server is running on port:", 8080)
	http.ListenAndServe(":8080", router)
}
