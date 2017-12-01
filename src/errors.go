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

// BadRequest responds with error status code 400, Bad Request.
func BadRequest(c *gin.Context) {
	m := "No routers match the request URL - https://api.cellosaur.us" + c.Request.URL.Path
	err := gin.H{"error": gin.H{"code": http.StatusBadRequest,
		"status": "Bad Request", "message": m}}
	c.JSON(http.StatusBadRequest, err)
}

// NotFound responds with error status code 404, Not Found.
func NotFound(c *gin.Context) {
	m := "The requested URI - https://api.cellosaur.us" + c.Request.URL.Path + " - does not match any resource."
	err := gin.H{"error": gin.H{"code": http.StatusNotFound,
		"status": "Not Found", "message": m}}
	c.JSON(http.StatusNotFound, err)
}

// InternalServerError responds with error status code 500, Internal Server Error.
func InternalServerError(c *gin.Context) {
	m := "Retry request. If error persists, open an issue on github."
	err := gin.H{"error": gin.H{"code": http.StatusInternalServerError,
		"status": "Internal Server Error", "message": m}}
	c.JSON(http.StatusInternalServerError, err)
}
