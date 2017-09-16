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
	err := gin.H{
		"error":   "Bad Request",
		"message": "No routers match the request URL: https://cellosaur.us" + c.Request.URL.Path,
	}
	c.JSON(http.StatusBadRequest, err)
}

// NotFound responds with error status code 404, Not Found.
func NotFound(c *gin.Context) {
	err := gin.H{
		"error":   "Not Found",
		"message": "The requested URI: https://cellosaur.us" + c.Request.URL.Path + " does not represent any resource on server",
	}
	c.JSON(http.StatusNotFound, err)
}

// InternalServerError responds with error status code 500, Internal Server Error.
func InternalServerError(c *gin.Context) {
	err := gin.H{
		"error":      "Internal Server Error",
		"message":    "An internal error has occurred in server, and steps are being taken to resolve issue.",
		"suggestion": "Try making the request after a few hours. If error persists, open an issue on github.",
	}
	c.JSON(http.StatusInternalServerError, err)
}
