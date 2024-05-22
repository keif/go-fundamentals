package models

import (
	"example.com/rest-api/db"
	"time"
)

type Event struct {
	ID          int64     `xorm:"pk autoincr"`
	Name        string    `xorm:"VARCHAR(255)" binding:"required"`
	Description string    `xorm:"TEXT" binding:"required"`
	Location    string    `xorm:"TEXT" binding:"required"`
	DateTime    time.Time `xorm:"TEXT" binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) AfterDelete() {}
func (e *Event) AfterFind()   {}
func (e *Event) AfterInsert() {}
func (e *Event) AfterQuery()  {}
func (e *Event) AfterUpdate() {}

func (e *Event) BeforeDelete() {}
func (e *Event) BeforeInsert() {}
func (e *Event) BeforeUpdate() {}

func GetAllEvents() (events []Event) {
	return events
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
    `
	// prepare for insertion - no sql attacks
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}
