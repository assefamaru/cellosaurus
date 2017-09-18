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

	router.LoadHTMLGlob("./src/cellosaurus/public/*")

	// Static resources
	router.StaticFile("/styles.css", "./src/cellosaurus/public/styles.css")
	router.StaticFile("/favicon.ico", "./src/cellosaurus/public/favicon.ico")

	// Handle root route
	router.GET("/", HomePage)

	// Handle api routes
	v1 := router.Group("/v1")
	for _, route := range routes {
		v1.Handle(route.Method, route.Endpoint, route.Handler)
	}

	// If no routers match the request url,
	// return 400 (Bad Request)
	router.NoRoute(BadRequest)

	// Listen and serve
	router.Run(":" + ctx.Port)
}
