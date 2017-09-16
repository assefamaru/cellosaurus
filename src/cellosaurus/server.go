package cellosaurus

import "github.com/gin-gonic/gin"

// Context is an api configuration containing
// router mode and server port information.
type Context struct {
	Mode string
	Port string
}

// Init server.
func Init(ctx *Context) {
	// Set database credentials
	SetDB()

	// Set gin mode
	gin.SetMode(ctx.Mode)

	// Gin router with default middleware: logger and recovery
	router := gin.Default()

	// Handle root route
	router.GET("/")

	// Handle api routes
	api := router.Group("/api")
	api.GET("/")
	for _, route := range routes {
		api.Handle(route.Method, route.Endpoint, route.Handler)
	}

	// If no routers match the request url,
	// return 400 (Bad Request)
	router.NoRoute(BadRequest)

	// Listen and serve
	router.Run(":" + ctx.Port)
}
