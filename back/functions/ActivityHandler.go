package functions

import (
	"database/sql"
	"log"

	// importing and init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const CreateUserStorageStatement = `CREATE TABLE IF NOT EXISTS activities
 (id NOT NULL AUTO_INCREMENT PRIMARY KEY, teamId INT, name VARCHAR(20), value INT)
 ENGINE=INNODB;`
const AddActivity = `INSERT INTO activities (teamId, name, value) VALUES(?,?,?);`
const UpdateActivity = `UPDATE activities SET value = ? WHERE teamId = ? AND id = ?;`

type ActivityHandler struct {
	databaseUrl        string
	databaseActivities *sql.DB
}

func NewActivityHandler(databaseUrl string) *ActivityHandler {
	handler := &ActivityHandler{}
	handler.databaseUrl = databaseUrl
	if err := handler.OpenConnection(); err != nil {
		return nil
	}
	if err := handler.InitTable(); err != nil {
		return nil
	}
	return handler
}

func (handler *ActivityHandler) ChangeActivity(teamId uint32, newValue int32, id uint32) error {
	stmt, err := handler.databaseActivities.Prepare(UpdateActivity)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(teamId, newValue, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (handler *ActivityHandler) AddActivity(teamId uint32, newValue int32, name string) error {
	stmt, err := handler.databaseActivities.Prepare(AddActivity)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(teamId, newValue, name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (handler *ActivityHandler) InitTable() error {
	stmt, err := handler.databaseActivities.Prepare(CreateUserStorageStatement)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (handler *ActivityHandler) OpenConnection() error {
	database, err := sql.Open("mysql", handler.databaseUrl)
	if err != nil {
		log.Println("Error occurred while creating database")
		return err
	} else {
		log.Println("Connection opened to DB")
	}
	handler.databaseActivities = database
	return nil
}

