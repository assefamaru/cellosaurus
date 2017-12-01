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
	err := gin.H{"error": gin.H{"code": http.StatusBadRequest, "status": "Bad Request", "message": m}}
	c.JSON(http.StatusBadRequest, err)
}
