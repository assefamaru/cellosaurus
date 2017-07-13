package main

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"os"

	raven "github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Sentry DSN for internal error logging.
func init() {
	raven.SetDSN("https://36b98457994b46efb1dea6c9ffd9eb70:19a5e80e08e043aeb6ef9f60693bbcf9@sentry.io/156124")
}

// Prepare database abstraction for later use.
func initDB() (*sql.DB, error) {
	user := os.Getenv("mysql_user")
	passwd := os.Getenv("mysql_passwd")
	mysqlDB := os.Getenv("mysql_cellosaurus_db")
	cred := user + ":" + passwd + "@tcp(127.0.0.1:3306)/" + mysqlDB
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
