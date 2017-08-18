package cellosaurus

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // go mysql driver
)

// DBAuthInfo contains local mysql username, password, and database name.
type DBAuthInfo struct {
	User string // local mysql user
	Pass string // local mysql password
	Name string // database name
	Host string // mysql host
}

// DB is a global datastore for database connection information.
var DB DBAuthInfo

// SetDB updates DB using environment settings.
func SetDB() {
	if DB.User = os.Getenv("CELL_DB_USER"); DB.User == "" {
		panic("Missing environment variable: CELL_DB_USER")
	}
	if DB.Pass = os.Getenv("CELL_DB_PASS"); DB.Pass == "" {
		panic("Missing environment variable: CELL_DB_PASS")
	}
	if DB.Name = os.Getenv("CELL_DB_NAME"); DB.Name == "" {
		panic("Missing environment variable: CELL_DB_NAME")
	}
	if DB.Host = os.Getenv("CELL_DB_HOST"); DB.Host == "" {
		panic("Missing environment variable: CELL_DB_HOST")
	}
}

// Database creates a new database connection.
func Database() (*sql.DB, error) {
	cred := DB.User + ":" + DB.Pass + "@tcp(" + DB.Host + ":3306)/" + DB.Name
	db, _ := sql.Open("mysql", cred)
	err := db.Ping()
	if err != nil {
		LogSentry(err)
	}
	return db, err
}
