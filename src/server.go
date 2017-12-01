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

	// Handle api routes
	v1 := router.Group("/v1")
	for _, route := range routes {
		v1.Handle(route.Method, route.Endpoint, route.Handler)
	}

	// If no routers match the request url,
	// return 400 (Bad Request)
	router.NoRoute(BadRequest)

	// Listen and serve
	router.Run(":" + c.Port)
}
