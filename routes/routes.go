package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luizflorentino/what-a-scheduler/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/schedule", controllers.Schedule)
	r.Run(":8000")
}