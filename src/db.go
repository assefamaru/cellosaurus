package cellosaurus

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// go mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type mysqlConfig struct {
	// MySQL user
	Username string

	// MySQL password
	Password string

	// MySQL database name
	Database string

	// Host of the MySQL instance.
	Host string

	// Port of the MySQL instance.
	Port string
}

var mysqlConf mysqlConfig

// SetMysqlConf updates mysqlConf using environment variables.
func SetMysqlConf() {
	if mysqlConf.Username = os.Getenv("cellosaurus_user"); mysqlConf.Username == "" {
		log.Fatal("cellosaurus_user environment variable must be set")
	}
	if mysqlConf.Password = os.Getenv("cellosaurus_pass"); mysqlConf.Password == "" {
		log.Fatal("cellosaurus_pass environment variable must be set")
	}
	if mysqlConf.Database = os.Getenv("cellosaurus_db"); mysqlConf.Database == "" {
		log.Fatal("cellosaurus_db environment variable must be set")
	}
	if mysqlConf.Host = os.Getenv("cellosaurus_host"); mysqlConf.Host == "" {
		log.Fatal("cellosaurus_host environment variable must be set")
	}
	mysqlConf.Port = "3306"
}

// dataSourceName returns a connection string suitable for sql.Open.
func (c mysqlConfig) dataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
}

// Database creates a new database connection.
func Database() (*sql.DB, error) {
	db, err := sql.Open("mysql", mysqlConf.dataSourceName())
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("mysql: could not establish a good connection: %v", err)
	}

	return db, nil
}