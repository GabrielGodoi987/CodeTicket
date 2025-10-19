package domain

import "errors"

var (
	ErrTicketPriceZero = errors.New("ticket price must be greater than zero")
)


type TicketType string

const (
	TicketTypeHalf TicketType = "half"
	TicketTypeFull TicketType = "full"
)

type Ticket struct {
	ID         string
	EventId    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}


func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (t *Ticket) CalculatePrice(){
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}


func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}

	return nil
}