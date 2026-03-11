package models

import (
	"time"

	"go-rest-backend/db"
)

type Event struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
}

func (e *Event) Save() error {

	query := `
	INSERT INTO events 
	(user_id, title, description, location, date_time)
	VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		e.UserID,
		e.Title,
		e.Description,
		e.Location,
		e.DateTime,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	e.ID = int(id)

	return nil
}

func GetAllEvents() ([]Event, error) {

	query := `
	SELECT id, user_id, title, description, location, date_time 
	FROM events
	`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {

		var event Event

		err := rows.Scan(
			&event.ID,
			&event.UserID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.DateTime,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}