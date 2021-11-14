package restapi

import (
	"net/http"
	"strconv"

	"github.com/GregersSR/taskinator/db"
	"github.com/gin-gonic/gin"
)

type addTaskRequest struct {
	Title  string `form:"title" json:"title" xml:"title"  binding:"required"`
	Device int64  `form:"device" json:"device" xml:"device"  binding:"required"`
	State  int16  `form:"state" json:"state" xml:"state"  binding:"required"`
}

func (r addTaskRequest) toDBType() db.NewTask {
	return db.NewTask{
		Title:  r.Title,
		Device: r.Device,
		State:  r.State,
	}
}

func addTask(c *gin.Context) {
	var request addTaskRequest
	c.Bind(&request)
	id, err := db.InsertTask(request.toDBType())
	if err != nil {
		c.String(http.StatusInternalServerError, "Error when inserting into db: %s", err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func deleteTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Could not parse %s as an integer", c.Param("id"))
		return
	}
	err = db.DeleteTask(id, true)
	if err != nil {

	}
}
