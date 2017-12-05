package cellosaurus

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReleaseInfo returns release information for current version of database.
func ReleaseInfo(c *gin.Context) {
	var rel Release
	if err := rel.Create(); err != nil {
		InternalServerError(c)
		return
	}
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "false"))
	Render(c, indent, rel)
}

// ListTerminologies returns a list of terminologies used in database.
func ListTerminologies(c *gin.Context) {
	var terms Terminologies
	if err := terms.List(); err != nil {
		InternalServerError(c)
		return
	}
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "false"))
	Render(c, indent, terms)
}

// ListCells handles GET requests for /cell-lines.
// func ListCells(c *gin.Context) {
// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "30"))
// 	all, _ := strconv.ParseBool(c.DefaultQuery("all", "false"))
// 	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "false"))
//
// 	total, err := Count()
// 	if err != nil {
// 		InternalServerError(c)
// 		return
// 	}
//
// 	if all {
// 		perPage = total
// 	}
//
// 	var cells Cells
// 	if err := cells.List(page, perPage); err != nil {
// 		InternalServerError(c)
// 		return
// 	}
//
// 	RenderWithMeta(c, page, perPage, total, indent, cells)
// }
