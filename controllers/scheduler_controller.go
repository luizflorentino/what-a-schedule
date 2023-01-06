package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizflorentino/what-a-scheduler/models"
)

func Schedule(c *gin.Context) {
	var schedule models.Schedule
	if c.ShouldBind(&schedule) != nil {
		c.JSON(200, gin.H{
			"Message":"Erro no binding",
		})
		return
	}

	url:= "https://wa.me/" + schedule.To + "?text=" + schedule.Message
	http.Redirect(c.Writer, c.Request, url, http.StatusPermanentRedirect)

	//TODO utilizar o sellenium para clicar nos botões
}
