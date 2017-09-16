package cellosaurus

import "github.com/gin-gonic/gin"

// TestFunc ...
func TestFunc(c *gin.Context) {
	InternalServerError(c)
}
