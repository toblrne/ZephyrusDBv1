package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toblrne/ZephyrusDB/db"
)

type Handler struct {
	dbDriver *db.Driver
}

func NewHandler(driver *db.Driver) *Handler {
	return &Handler{
		dbDriver: driver,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.dbDriver.Write("users", user.Name, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) GetUser(c *gin.Context) {
	name := c.Param("name")
	var user db.User
	if err := h.dbDriver.Read("users", name, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// In your handler.go

func (h *Handler) UpdateUser(c *gin.Context) {
	name := c.Param("name")
	var userUpdates db.User

	if err := c.ShouldBindJSON(&userUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user db.User
	if err := h.dbDriver.Read("users", name, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Assume the user's name cannot be changed, or add logic to handle name changes
	if err := h.dbDriver.Write("users", name, &userUpdates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userUpdates)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	name := c.Param("name")
	if err := h.dbDriver.Delete("users", name); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or could not be deleted"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.dbDriver.ReadAll("users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
