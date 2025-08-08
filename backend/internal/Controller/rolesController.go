package Controller

import "github.com/gin-gonic/gin"

func GetAllRoles(c *gin.Context) {
	roles := []string{"admin", "user", "guest"}
	c.JSON(200, gin.H{"roles:": roles})
}
