package restapi

import (
	"net/http"
	"strconv"

	"github.com/GregersSR/taskinator/db"
	"github.com/gin-gonic/gin"
)

type addUserRequest struct {
	Name   string `form:"name" json:"name" xml:"name"  binding:"required"`
	Email  string `form:"email" json:"email" xml:"email"  binding:"required"`
	Admin  bool   `form:"admin" json:"admin" xml:"admin"`
	Token  string `form:"token" json:"token" xml:"token"  binding:"required"`
	Active bool   `form:"active" json:"active" xml:"active"`
}

func (r addUserRequest) toDBType() db.CreateUserDTO {
	return db.CreateUserDTO{
		Name:   r.Name,
		Email:  r.Email,
		Admin:  r.Admin,
		Token:  r.Token,
		Active: r.Active,
	}
}

func addUser(c *gin.Context) {
	var request addUserRequest
	err := c.Bind(&request)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request: %v", err)
	}
	id, err := db.InsertUser(request.toDBType())
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}
	c.JSON(http.StatusCreated, id)
}

func deleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Could not parse %s as an integer", c.Param("id"))
		return
	}
	db.DeleteUser(id, true)
}

func getUserDetails(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Could not parse %s as an integer", c.Param("id"))
		return
	}
	user, err := db.GetUser(id)
	if err != nil {
		c.String(http.StatusNotFound, "%v", err)
		return
	}
	c.JSON(http.StatusOK, user)
}
