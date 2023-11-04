package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/user", handler.CreateUser)
	router.GET("/user/:name", handler.GetUser)
	router.PUT("/user/:name", handler.UpdateUser)
	router.DELETE("/user/:name", handler.DeleteUser)
	router.GET("/users", handler.ListUsers)

	return router
}
