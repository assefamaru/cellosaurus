package main

import (
	"database/sql"
	"os"

	raven "github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
)

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
