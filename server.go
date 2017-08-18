package cellosaurus

import "github.com/gin-gonic/gin"

// Init server.
func Init(ctx *Context) {
	// Set database credentials
	SetDB()

	SetMode(ctx.Mode)
	SetPort(ctx.Port)
	SetVersion(ctx.Version)

	// Gin router with default middleware: logger and recovery
	router := gin.Default()

	// Serve favicon
	router.StaticFile("/favicon.ico", "./static/images/favicon.ico")

	router.GET("/", RootHandler)

	v := router.Group(Version() + "/")
	v.GET("/", RootHandler)
	for _, route := range routes {
		v.Handle(route.Method, route.Endpoint, route.Handler)
	}

	// If no routers match the request url, return 400 (Bad Request)
	router.NoRoute(func(c *gin.Context) {
		BadRequest(c, "The endpoint "+c.Request.URL.Path+" is not well formed")
	})

	// Listen and serve on config port
	router.Run(":" + Port())
}
