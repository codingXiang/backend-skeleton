package delivery

import "github.com/gin-gonic/gin"

//HttpHandler http流量 handler
type HttpHandler interface {
	CreateDepartment(c *gin.Context) error
	CreateUser(c *gin.Context) error
	GetDepartment(c *gin.Context) error
	GetUser(c *gin.Context) error
	ModifyDepartment(c *gin.Context) error
	ModifyUser(c *gin.Context) error
	DeleteDepartment(c *gin.Context) error
	DeleteUser(c *gin.Context) error
}
