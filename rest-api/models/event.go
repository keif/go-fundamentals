package models

import (
	"errors"
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

func (e *Event) CancelRegistration(userId int64) error {
	var err error
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func (e *Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
    `
	// prepare for insertion
	// 100% optional
	// can lead to better performance if the sql is executed multiple times, potentially with different data
	// ONLY TRUE if the connection is not closed
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Exec is since we're changing data
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name=?, description=?, location=?, dateTime=?, user_id=?
    WHERE id=?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) // it's a select, no need to exec
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	if id <= 0 {
		return nil, errors.New("Invalid id")
	}
	query := "SELECT * FROM events WHERE id=?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
