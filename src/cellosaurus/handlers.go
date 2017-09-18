package cellosaurus

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IndexCells returns a list of cell lines.
// Handles GET requests for /cell_lines.
func IndexCells(c *gin.Context) {
	var cells Cells
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "30"))
	all, _ := strconv.ParseBool(c.DefaultQuery("all", "false"))
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))
	include := c.Query("include")
	if all {
		err := cells.List()
		if err != nil {
			InternalServerError(c)
			return
		}
		c.Writer.Header().Set("Total-Records", strconv.Itoa(len(cells)))
		RenderJSON(c, indent, cells)
	} else {
		err := cells.ListPaginated(page, limit)
		if err != nil {
			InternalServerError(c)
			return
		}
		total, err := Count("cellosaurus")
		if err != nil {
			InternalServerError(c)
			return
		}
		WriteHeader(c, "/cell_lines", page, limit, total)
		RenderJSONwithMeta(c, indent, page, limit, total, include, cells)
	}
}

// ShowCell returns a single cell line using id/name.
// Handles GET requests for /cell_lines/{id}.
func ShowCell(c *gin.Context) {
	var cell Cell
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))
	typ := c.DefaultQuery("type", "accession")
	id := c.Param("id")
	err := cell.Find(id, typ)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFound(c)
		} else if err.Error() == "Unknown type" {
			BadRequest(c)
		} else {
			InternalServerError(c)
		}
		return
	}
	RenderJSON(c, indent, cell)
}

// SearchCell finds a cell line using either accession, identifier, or synonyms.
func SearchCell(c *gin.Context) {
	var cell Cell
	indent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))
	id := c.Param("id")
	err := cell.Find(id, "accession")
	if err == nil {
		RenderJSON(c, indent, cell)
		return
	} else if err != sql.ErrNoRows {
		InternalServerError(c)
		return
	}
	err = cell.Find(id, "identifier")
	if err == nil {
		RenderJSON(c, indent, cell)
		return
	} else if err != sql.ErrNoRows {
		InternalServerError(c)
		return
	}
	err = cell.Find(id, "synonym")
	if err == nil {
		RenderJSON(c, indent, cell)
	} else if err != sql.ErrNoRows {
		InternalServerError(c)
	} else {
		NotFound(c)
	}
}

// HomePage renders home page html for root route.
func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Cellosaurus API",
	})
}
