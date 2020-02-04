package cellosaurus

import (
	"net/http"

	raven "github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
)

// Sentry DSN for internal error logging.
func init() {
	raven.SetDSN("https://36b98457994b46efb1dea6c9ffd9eb70:19a5e80e08e043aeb6ef9f60693bbcf9@sentry.io/156124")
}

// logSentry submits internal errors to Sentry.
func logSentry(err error) {
	raven.CaptureError(err, nil)
}

// Error is a custom error structure.
type Error struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// renderError renders a custom error.
func renderError(c *gin.Context, indent bool, err Error) {
	if indent {
		c.IndentedJSON(err.Code, gin.H{"error": err})
	} else {
		c.JSON(err.Code, gin.H{"error": err})
	}
}

// BadRequest responds with error status code 400, Bad Request.
func BadRequest(c *gin.Context) {
	var err Error
	err.Code = http.StatusBadRequest
	err.Status = "Bad Request"
	err.Message = "No routers match the request URL - https://api.cellosaur.us" + c.Request.URL.Path
	renderError(c, false, err)
}

// NotFound responds with error status code 404, Not Found.
func NotFound(c *gin.Context) {
	var err Error
	err.Code = http.StatusNotFound
	err.Status = "Not Found"
	m := "The requested URI - https://api.cellosaur.us" + c.Request.URL.Path
	n := " - does not match any resource in database."
	err.Message = m + n
	renderError(c, false, err)
}

// InternalServerError responds with error status code 500, Internal Server Error.
func InternalServerError(c *gin.Context) {
	var err Error
	err.Code = http.StatusInternalServerError
	err.Status = "Internal Server Error"
	err.Message = "Retry request. If error persists, open an issue on github."
	renderError(c, false, err)
}