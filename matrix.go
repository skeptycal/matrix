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

// getEnvConnectionString - get environment variable from mySqlUserVariable
func getEnvConnectionString(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("environment variable <%v> not found", key)
	}
	return value, nil
}

// Check performs a connection check on the mysql database connection
func Check() {
	mysql_username, err := getEnvConnectionString(mySqlUserVariable)
	if err != nil {
		log.Info(err)
	}

	log.Info("mysql username: ", mysql_username)

	Db, err := sql.Open("mysql", mysql_username)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

}

// Notes:
/*
// DSN (Data Source Name)
// The Data Source Name has a common format, like e.g. PEAR DB uses it, but without type-prefix (optional parts marked by squared brackets):
*/
