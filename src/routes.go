package cellosaurus

import "github.com/gin-gonic/gin"

// HTTP methods.
const (
	GET    string = "GET"
	HEAD   string = "HEAD"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
	OPTION string = "OPTION"
	PATCH  string = "PATCH"
)

// Route is a routing model.
type Route struct {
	Method   string
	Endpoint string
	Handler  gin.HandlerFunc
}

// Routes is a collection of Route.
type Routes []Route

var routes = Routes{
	Route{GET, "/release-info", GetReleaseInfo},
	Route{GET, "/terminologies", ListTerminologies},
	Route{GET, "/references", ListReferences},
	Route{GET, "/cell-lines", ListCells},
	Route{GET, "/cell-lines/:id", FindCell},
}