package database

import (
	"database/sql"
	"fmt"

	// importing nad init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DBConnector connects to DB
type DBConnector struct {
	databaseActivities *sql.DB
}

// Open database connection
func (conn *DBConnector) Open(path string) {
	database, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println("Error occurred while creating database")
		return
	}
	conn.databaseActivities = database
}

// CreateActivitiesDatabaseIfDoesNotExist creates database for activities if does not exist
func (conn *DBConnector) CreateActivitiesDatabaseIfDoesNotExist() bool {
	statement, err := conn.databaseActivities.Prepare(`CREATE TABLE IF NOT EXISTS Activities
	(id varchar(36), name TEXT, orgId varchar(36), PRIMARY KEY (id))`)
	if err != nil {
		fmt.Println("Error occurred while creating database ", err)
		return false
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Println("Error occurred while exec creation of database ", err)
		return false
	}

	return true
}