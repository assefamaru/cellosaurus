package cellosaurus

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // go mysql driver
)

// DBAuthInfo contains local mysql username,
// password, database name, and host information.
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
	if DB.User = os.Getenv("CDB_USER"); DB.User == "" {
		log.Fatal("CDB_USER must be set")
	}
	if DB.Pass = os.Getenv("CDB_PASS"); DB.Pass == "" {
		log.Fatal("CDB_PASS must be set")
	}
	if DB.Name = os.Getenv("CDB_NAME"); DB.Name == "" {
		log.Fatal("CDB_NAME must be set")
	}
	if DB.Host = os.Getenv("CDB_HOST"); DB.Host == "" {
		log.Fatal("CDB_HOST must be set")
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
