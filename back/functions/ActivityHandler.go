package functions

import (
	"database/sql"
	"log"

	// importing and init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const CreateUserStorageStatement = `CREATE TABLE IF NOT EXISTS activities
 (id INT NOT NULL AUTO_INCREMENT, teamId INT, name VARCHAR(20), value INT, PRIMARY KEY (id))
 ENGINE=INNODB;`
const AddActivity = `INSERT INTO activities (teamId, name, value) VALUES(?,?,0);`
const UpdateActivity = `UPDATE activities SET value = ? WHERE teamId = ? AND id = ?;`
const GetActivitiesForTeam = `SELECT id, name, value FROM activities WHERE teamId = ?;`

type ActivityHandler struct {
	databaseUrl        string
	databaseActivities *sql.DB
}

type ActivityData struct {
	id    uint32
	name  string
	value uint32
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
	_, err = stmt.Exec(newValue, teamId, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (handler *ActivityHandler) AddActivity(teamId uint32, name string) error {
	stmt, err := handler.databaseActivities.Prepare(AddActivity)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(teamId, name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (handler *ActivityHandler) GetActivitiesForTeam(teamId uint32) ([]ActivityData, error) {
	stmt, err := handler.databaseActivities.Prepare(GetActivitiesForTeam)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := stmt.Query(teamId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var activities []ActivityData
	for rows.Next() {
		activity := &ActivityData{}
		err := rows.Scan(&activity.id, &activity.name, &activity.value)
		if err != nil {
			log.Println(err)
		}
		activities = append(activities, *activity)
	}

	return activities, nil
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
	log.Println("Created activities table")
	return nil
}

func (handler *ActivityHandler) OpenConnection() error {
	database, err := sql.Open("mysql", handler.databaseUrl)
	if err != nil {
		log.Println("Error occurred while opening database")
		return err
	} else {
		log.Println("Connection opened to DB")
	}
	handler.databaseActivities = database
	return nil
}

