package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrTicketPriceZero = errors.New("ticket price must be greater than zero")
)

type TicketKind string

const (
	TicketTypeHalf TicketKind = "half"
	TicketTypeFull TicketKind = "full"
)

// Errors
var (
	ErrInvalidTicketKind = errors.New("invalid ticket kind")
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

func NewTicket(event *Event, spot *Spot, ticketKind TicketKind) (*Ticket, error) {
	if !IsValidTicketKind(ticketKind) {
		return nil, ErrInvalidTicketKind
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketKind: ticketKind,
		Price:      event.Price,
	}
	
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}

	return ticket, nil
}

func IsValidTicketKind(ticketKind TicketKind) bool {
	return ticketKind == TicketTypeHalf || ticketKind == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}

	return nil
}
