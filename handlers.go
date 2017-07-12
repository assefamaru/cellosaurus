package main

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"

	raven "github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

// Sentry DSN for internal error logging.
func init() {
	raven.SetDSN("https://36b98457994b46efb1dea6c9ffd9eb70:19a5e80e08e043aeb6ef9f60693bbcf9@sentry.io/156124")
}

// Prepare database abstraction for later use.
func initDB() (*sql.DB, error) {
	cred := os.Getenv("mysql_user") + ":" + os.Getenv("mysql_passwd") + "@tcp(127.0.0.1:3306)/cellosaurus"
	db, err := sql.Open("mysql", cred)
	if err != nil {
		raven.CaptureError(err, nil)
	}
	// Establish and test database connection.
	err = db.Ping()
	if err != nil {
		raven.CaptureError(err, nil)
	}
	return db, err
}

// Handle error messages (all except no route match errors).
func handleError(c *gin.Context, err error, code int, message string) {
	if err != nil {
		raven.CaptureError(err, nil)
	}
	c.JSON(code, gin.H{
		"error": gin.H{
			"code":    code,
			"message": message,
		},
	})
}

// writeHeaderLinks writes pagination links in response header.
// Links available under 'Link' header, including (prev, next, first, last).
func writeHeaderLinks(c *gin.Context, endpoint string, page int, total int, limit int) {
	var (
		prev    string
		prevRel string
		next    string
		nextRel string
	)
	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	first := fmt.Sprintf("<https://cellosaurus.pharmacodb.com/v1%s?page=%d&per_page=%d>", endpoint, 1, limit)
	if (page > 1) && (page <= lastPage) {
		prev = fmt.Sprintf("<https://cellosaurus.pharmacodb.com/v1%s?page=%d&per_page=%d>", endpoint, page-1, limit)
		prevRel = "; rel=\"prev\", "
	}
	if (page >= 1) && (page < lastPage) {
		next = fmt.Sprintf("<https://cellosaurus.pharmacodb.com/v1%s?page=%d&per_page=%d>", endpoint, page+1, limit)
		nextRel = "; rel=\"next\", "
	}
	last := fmt.Sprintf("<https://cellosaurus.pharmacodb.com/v1%s?page=%d&per_page=%d>", endpoint, lastPage, limit)

	linknp := prev + prevRel + next + nextRel
	linkfl := first + "; rel=\"first\", " + last + "; rel=\"last\""
	link := linknp + linkfl

	c.Writer.Header().Set("Link", link)
}

// APIVersionEndpoint returns a welcome string for root versioned endpoints.
func APIVersionEndpoint(c *gin.Context) {
	message := "Welcome to Cellosaurus API, see: https://github.com/assefamaru/cellosaurus-api"
	c.String(http.StatusOK, message)
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
