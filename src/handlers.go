package cellosaurus

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListCells handles GET requests for /cell-lines.
func ListCells(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "30"))
	all, _ := strconv.ParseBool(c.DefaultQuery("all", "false"))
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "false"))

	total, err := Count()
	if err != nil {
		InternalServerError(c)
		return
	}

	if all {
		perPage = total
	}

	var cells Cells
	if err := cells.List(page, perPage); err != nil {
		InternalServerError(c)
		return
	}

	RenderWithMeta(c, page, perPage, total, indent, cells)
}
