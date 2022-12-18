// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/KnightHacks/knighthacks_shared/models"
)

type Connection interface {
	IsConnection()
}

type Event struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
}

func (Event) IsEntity() {}

type EventsConnection struct {
	TotalCount int              `json:"totalCount"`
	PageInfo   *models.PageInfo `json:"pageInfo"`
	Events     []*Event         `json:"events"`
}

func (EventsConnection) IsConnection() {}

type NewEvent struct {
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	HackathonID string    `json:"hackathonId"`
}

type UpdatedEvent struct {
	Name        *string    `json:"name"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Description *string    `json:"description"`
	Location    *string    `json:"location"`
}
