package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//initialize router
	r := gin.Default()

	//initialize routes
	initializeRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080

}
