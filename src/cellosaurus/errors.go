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

// LogSentry submits private errors to sentry.
func LogSentry(err error) {
	raven.CaptureError(err, nil)
}

// BadRequest responds with error status code 400, Bad Request.
func BadRequest(c *gin.Context) {
	message := "Bad Request. No routers match the request URL: https://cellosaur.us" + c.Request.URL.Path
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"error": gin.H{
			"status":  http.StatusBadRequest,
			"message": message,
		},
	})
}

// NotFound responds with error status code 404, Not Found.
func NotFound(c *gin.Context) {
	message := "Not Found. The requested URI: https://cellosaur.us" + c.Request.URL.Path + " does not represent any resource on server"
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"error": gin.H{
			"status":  http.StatusNotFound,
			"message": message,
		},
	})
}

// InternalServerError responds with error status code 500, Internal Server Error.
func InternalServerError(c *gin.Context) {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"error": gin.H{
			"status":     http.StatusInternalServerError,
			"message":    "Internal Server Error",
			"suggestion": "Retry request after a few hours. If error persists, open an issue on github.",
		},
	})
}
