package cellosaurus

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Render renders JSON response.
func Render(c *gin.Context, indent bool, obj interface{}) {
	if indent {
		c.IndentedJSON(http.StatusOK, obj)
	} else {
		c.JSON(http.StatusOK, obj)
	}
}

// RenderWithMeta renders JSON response with metadata.
func RenderWithMeta(c *gin.Context, page int, perPage int, total int, indent bool, data interface{}) {
	var obj interface{}

	if perPage == total {
		obj = gin.H{"data": data}
	} else {
		lastPage := int(math.Ceil(float64(total) / float64(perPage)))
		meta := gin.H{"first-page": 1, "last-page": lastPage,
			"current-page": page, "per-page": perPage, "total-cell-lines": total}
		obj = gin.H{"meta": meta, "data": data}
	}

	Render(c, indent, obj)
}
