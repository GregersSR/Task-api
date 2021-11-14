package restapi

import (
	"net/http"
	"strconv"

	"github.com/GregersSR/taskinator/db"
	"github.com/gin-gonic/gin"
)

type addDeviceRequest struct {
	Name        string `form:"name" json:"name" xml:"name"  binding:"required"`
	Description string `form:"description" json:"description" xml:"description"  binding:"required"`
	Token       string `form:"token" json:"token" xml:"token"  binding:"required"`
}

func (r addDeviceRequest) toDBType() db.NewDevice {
	return db.NewDevice{
		Name:        r.Name,
		Description: r.Description,
		Token:       r.Token,
	}
}

func addDevice(c *gin.Context) {
	var request addDeviceRequest
	c.Bind(&request)
	id, err := db.InsertDevice(request.toDBType())
	if err != nil {
		c.String(http.StatusInternalServerError, "Error when inserting into db: %s", err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func deleteDevice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Could not parse %s as an integer", c.Param("id"))
		return
	}
	err = db.DeleteDevice(id, true)
	if err != nil {

	}
}
