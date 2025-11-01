package main

import (
	"api/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()

	router.Use(CORSMiddleware())

	router.StaticFS("/assets", http.Dir("../spa/dist/assets/"))

	apiRoutes := router.Group("/api")
	{
		// Groceries
		apiRoutes.GET("/groceries/:householdId", routes.GetGroceries)
		apiRoutes.PUT("/groceries", routes.CreateGroceryItem)
		apiRoutes.POST("/groceries", routes.UpdateGroceryItem)
		apiRoutes.DELETE("/groceries/:householdId/:id", routes.DeleteGroceryItem)
		apiRoutes.POST("/groceries/batchDelete", routes.BatchDeleteGroceryItems)
		apiRoutes.POST("/groceries/magic", routes.GroceryMagic)
		apiRoutes.POST("/tasks/schedule", routes.ScheduleTask)

		// Households
		apiRoutes.PUT("/households", routes.CreateHousehold)
		apiRoutes.POST("/households/join/:householdId/:userId", routes.JoinHousehold)
		apiRoutes.POST("/households/leave/:householdId/:userId", routes.LeaveHousehold)

		// Users
		apiRoutes.PUT("/users", routes.CreateUser)
		apiRoutes.GET("/users/:id", routes.GetUser)
	}

	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.AbortWithStatus(404)
			return
		}

		c.File("../spa/dist/index.html")
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router.Run(":57457")
}
