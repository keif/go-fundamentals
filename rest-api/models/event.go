package models

import "time"

type Event struct {
	ID          int64     `xorm:"pk autoincr"`
	Name        string    `xorm:"VARCHAR(255)"`
	Description string    `xorm:"TEXT"`
	DateTime    time.Time `xorm:"TEXT"`
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

func (e *Event) Save() {
	// TODO: store to database
	events = append(events, *e)
}
