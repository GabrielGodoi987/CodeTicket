package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEventNameRequired = errors.New("Event name is required")
	ErrEventDateFuture   = errors.New("Event date must be in the future")
	ErrCapacityZero      = errors.New("event capacity must be greater than zero")
	ErrEventCapacityZero = errors.New("event price must be greater than zero")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	RatingL12   Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int    // de acordo com o id, vamos chamar o sistema externo diferente
	Spots        []Spot // estamos utilizando slice(array dinâmico que diz que o nosso evennt pode ter vários spots)
	Tickets      []Ticket
}

func NewEvent(name, location, organization string, rating Rating, date time.Time, capacity int, price float64, imageUrl string, partnerID int) (*Event, error) {
	event := &Event{
		ID:           uuid.New().String(),
		Name:         name,
		Location:     location,
		Organization: organization,
		Rating:       rating,
		Date:         date,
		Capacity:     capacity,
		Price:        price,
		ImageURL:     imageUrl,
		PartnerID:    partnerID,
		Spots:        make([]Spot, 0),
	}
	if err := event.Validate(); err != nil {
		return nil, err
	}
	return event, nil
}

func (e Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}
	if e.Date.Before(time.Now()) {
		return ErrEventDateFuture
	}

	if e.Capacity <= 0 {
		return ErrCapacityZero
	}

	if e.Price <= 0 {
		return ErrEventCapacityZero
	}

	return nil
}

func (e *Event) addSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)

	if err != nil {
		return nil, err
	}

	// adicionamos mais um spot na lista de spots
	e.Spots = append(e.Spots, *spot)

	return spot, nil
}
