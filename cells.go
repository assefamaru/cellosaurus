package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Accession is a unique accession id.
type Accession struct {
	Primary   string  `json:"primary"`
	Secondary *string `json:"secondary"`
}

// Cell is a cell line data type.
type Cell struct {
	ID         int       `json:"id"`
	Identifier string    `json:"identifier"`
	Accession  Accession `json:"accession"`
}

// IndexCell returns a list of all cell lines in database (paginated by default).
func IndexCell(c *gin.Context) {
	var (
		cell  Cell
		cells []Cell
	)

	db, err := initDB()
	defer db.Close()
	if err != nil {
		handleError(c, nil, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	shouldIndent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))

	if isTrue, _ := strconv.ParseBool(c.DefaultQuery("all", "false")); isTrue {
		rows, er := db.Query("SELECT id, identifier, accession, `as` FROM cells;")
		defer rows.Close()
		if er != nil {
			handleError(c, er, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		for rows.Next() {
			err = rows.Scan(&cell.ID, &cell.Identifier, &cell.Accession.Primary, &cell.Accession.Secondary)
			if err != nil {
				handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			cells = append(cells, cell)
		}
		if shouldIndent {
			c.IndentedJSON(http.StatusOK, gin.H{
				"data":        cells,
				"total":       len(cells),
				"description": "List of all cell lines in Cellosaurus",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data":        cells,
				"total":       len(cells),
				"description": "List of all cell lines in Cellosaurus",
			})
		}
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "30"))

	s := (page - 1) * limit
	SQL := "SELECT id, identifier, accession, `as` FROM cells"
	query := fmt.Sprintf("%s limit %d,%d", SQL, s, limit)
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	for rows.Next() {
		err = rows.Scan(&cell.ID, &cell.Identifier, &cell.Accession.Primary, &cell.Accession.Secondary)
		if err != nil {
			handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		cells = append(cells, cell)
	}
	row := db.QueryRow("SELECT COUNT(*) FROM cells;")
	var total int
	err = row.Scan(&total)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Write pagination links in response header.
	writeHeaderLinks(c, "/cell_lines", page, total, limit)

	if shouldIndent {
		c.IndentedJSON(http.StatusOK, gin.H{
			"data":        cells,
			"total":       total,
			"page":        page,
			"per_page":    limit,
			"description": "List of all cell lines in Cellosaurus",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":        cells,
			"total":       total,
			"page":        page,
			"per_page":    limit,
			"description": "List of all cell lines in Cellosaurus",
		})
	}
}

// ShowCell returns a cell line using ID, Identifier, Accession, or Synonyms.
func ShowCell(c *gin.Context) {

}
