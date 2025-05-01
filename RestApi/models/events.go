package models

import (
	"database/sql"
	"time"
)

var DB *sql.DB

type Event struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserId      int       `json:"userid"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, datetime, userid)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.Id)
}

func GetEvents() ([]Event, error) {
	rows, err := DB.Query("SELECT id, name, description, location, datetime, userid FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
		if err != nil {
			return nil, err // return if any row scan fails
		}
		events = append(events, e)
	}

	return events, nil
}
