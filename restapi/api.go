package restapi

import "github.com/gin-gonic/gin"

func Serve() {
	r := gin.Default()
	init_routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
