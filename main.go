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
	mySqlUserDefault  = ""
)

// DSN (Data Source Name)
// The Data Source Name has a common format, like e.g. PEAR DB uses it, but without type-prefix (optional parts marked by squared brackets):

// getEnv - get value from environment variable key
func getEnv(key string, defaultValue string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue, fmt.Errorf("environment variable <%v> not found", key)
	}
	return value, nil
}

func init() {

	mysql_username, err := getEnv(mySqlUserVariable, mySqlUserDefault)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("mysql username: ", mysql_username)

	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

}

// Notes:
/* Matrix - an organizational map based on the Eisenhower Matrix.

todo - list of steps to complete

// Data:
- structure to store actions
- structure to store calendar days
- structure to store resources
- structure to store ideas
- structure to store relationship graphs

// Major Parts:
- mysql database
- go link to mysql database
- frontend templates

// CRUD database
Create -

Read -

Update -

Delete -

// API funcionality

// API Documentation


*/
