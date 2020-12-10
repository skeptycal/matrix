package matrix

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	mySqlUserVariable = "MYSQL_USERNAME"
)

type dbConfig struct {
	userName string
	host     string
}

// getEnvConnectionString - get environment variable from mySqlUserVariable
func getEnvConnectionString(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable <%v> not found", key)
	}

	value = value + "/test"

	return value, nil
}

func dbConnect() (*sql.DB, error) {
	mysqlConnectionString, err := getEnvConnectionString(mySqlUserVariable)
	if err != nil {
		log.Error(err)
	}

	log.Info("mysql username: ", mysqlConnectionString)

	// Open database connection.
	db, err := sql.Open("mysql", mysqlConnectionString)
	if err != nil {
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

// Check performs a connection check on the mysql database connection
func Check() {
	db, err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}

	// defer the close until  the main function is done
	defer db.Close()

	// perform query test
	response, err := db.Query("SHOW DATABASES;")
	if err != nil {
		log.Fatal(err)
	}

	log.Info("MySQL query response: ", response)

}

// Notes:
/*
// DSN (Data Source Name)
// The Data Source Name has a common format, like e.g. PEAR DB uses it, but without type-prefix (optional parts marked by squared brackets):
*/
