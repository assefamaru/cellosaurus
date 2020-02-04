package cellosaurus

import "github.com/gin-gonic/gin"

// Context contains router mode
// and server port information.
type Context struct {
	Mode string
	Port string
}

// Init server.
func Init(c Context) {
	// Set database credentials
	SetMysqlConf()

	// Set gin mode
	gin.SetMode(c.Mode)

	// Gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.Static("/assets", "./static/img/*")
	router.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	// Handle api routes
	v33 := router.Group("/v33")
	for _, route := range routes {
		v33.Handle(route.Method, route.Endpoint, route.Handler)
	}

	// If no routers match the request url,
	// return 400 (Bad Request)
	router.NoRoute(BadRequest)

	// Listen and serve
	router.Run(":" + c.Port)
}