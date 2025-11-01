package routes

import (
	"api/models"
	"api/providers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ScheduleTask(c *gin.Context) {
	var request models.ScheduleTaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := providers.CreateTaskSchedule(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
