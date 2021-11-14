package restapi

import "github.com/gin-gonic/gin"

func init_routes(r *gin.Engine) {
	// Managing users
	//r.GET("/users", getUsers)
	r.POST("/users", addUser)
	r.GET("/users/:id", getUserDetails)
	r.DELETE("/users/:id", deleteUser)

	// Managing devices
	r.POST("/devices", addDevice)
	//r.GET("/devices", getDevices)
	//r.GET("/devices/:id", getDeviceDetails)
	r.DELETE("/devices/:id", deleteDevice)

	// Managing tasks and responses
	r.POST("/devices/:id/tasks", addTask)
	//r.GET("/devices/:id/tasks", getTasks)
	r.DELETE("/devices/:id/tasks/:taskid", deleteTask)
	//r.POST("/devices/:id/tasks/:taskid/execute", executeTask)
	//r.POST("/devices/:id/tasks/:taskid/response", addResponse)
	//r.GET("/devices/:id/tasks/:taskid/response", getResponse)
}
