package database

import "database/sql"

type EventsModel struct {
	DB *sql.DB
}

type Event struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"ownerId" binding:"required"`
	Name        string `json:"name" binding:"required, min=3"`
	Description string `json:"description" binding:"required,min10"`
	Date        string `json:"date" binding:"required, datetime=2007-01-02"`
	Location    string `json:"location" binding:"required, min=3"`
}
