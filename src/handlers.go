package cellosaurus

import (
	"database/sql"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Render applies indentation settings to response.
func Render(c *gin.Context, indent bool, obj interface{}) {
	if indent {
		c.IndentedJSON(http.StatusOK, obj)
	} else {
		c.JSON(http.StatusOK, obj)
	}
}

// GetReleaseInfo returns release information for current version of database.
func GetReleaseInfo(c *gin.Context) {
	var rel Release
	if err := rel.Create(); err != nil {
		InternalServerError(c)
		return
	}
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))
	Render(c, indent, rel)
}

// ListTerminologies returns a list of terminologies used in database.
func ListTerminologies(c *gin.Context) {
	var terms Terminologies
	if err := terms.List(); err != nil {
		InternalServerError(c)
		return
	}
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))
	Render(c, indent, terms)
}

// ListCells handles GET requests for /cell-lines.
func ListCells(c *gin.Context) {
	var cells Cells

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))

	total, err := totalCells()
	if err != nil {
		InternalServerError(c)
		return
	}

	lastPage := int(math.Ceil(float64(total) / float64(perPage)))

	// Set max number of cell lines to return per request
	if perPage > 100 {
		perPage = 100
	}

	cells.Meta.Page = page
	cells.Meta.PerPage = perPage
	cells.Meta.LastPage = lastPage
	cells.Meta.Total = total

	if err := cells.List(); err != nil {
		InternalServerError(c)
		return
	}

	Render(c, indent, cells)
}

// FindCell handles GET requests for /cell-lines/:id.
func FindCell(c *gin.Context) {
	var cell Cell
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))

	cell.ID = c.Param("id")
	err := cell.Find()
	if err != nil {
		if err == sql.ErrNoRows {
			NotFound(c)
		} else {
			InternalServerError(c)
		}
		return
	}

	Render(c, indent, cell)
}

// ListReferences handles GET requests for /references.
func ListReferences(c *gin.Context) {
	var refs References

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))

	total, err := totalRefs()
	if err != nil {
		InternalServerError(c)
		return
	}

	lastPage := int(math.Ceil(float64(total) / float64(perPage)))

	// Set max number of references to return per request
	if perPage > 100 {
		perPage = 100
	}

	refs.Meta.Page = page
	refs.Meta.PerPage = perPage
	refs.Meta.LastPage = lastPage
	refs.Meta.Total = total

	if err := refs.List(); err != nil {
		InternalServerError(c)
		return
	}

	Render(c, indent, refs)
}
