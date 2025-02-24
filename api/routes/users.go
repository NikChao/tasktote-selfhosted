package routes

import (
	"api/providers"
	db "api/proxy/sqlite"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := providers.CreateUser()

	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	database, _ := db.NewDB()
	defer database.Close()

	user, err := providers.GetOrCreateUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, user)
}
