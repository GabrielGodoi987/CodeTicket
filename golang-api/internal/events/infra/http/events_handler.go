package http

import (
	"encoding/json"
	"net/http"

	usecases "github.com/gabrielgodoi987/CodeTicket/golang-api/internal/events/useCases"
)

type EventsHandler struct {
	listEventsUseCase  *usecases.ListEventsUseCase
	getEventUseCase    *usecases.GetEventUseCase
	createEventUseCase *usecases.CreateEventUseCase
	buyTicketsUseCase  *usecases.BuyTicketsUseCase
	createSpotsUseCase *usecases.CreateSpotsUseCase
	listSpotsUseCase   *usecases.ListSpotsUseCase
}

func NewEventsHandler(
	listEventsUseCase *usecases.ListEventsUseCase,
	getEventUseCase *usecases.GetEventUseCase,
	createEventUseCase *usecases.CreateEventUseCase,
	buyTicketsUseCase *usecases.BuyTicketsUseCase,
	createSpotsUseCase *usecases.CreateSpotsUseCase,
	listSpotsUseCase *usecases.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase:  listEventsUseCase,
		getEventUseCase:    getEventUseCase,
		createEventUseCase: createEventUseCase,
		buyTicketsUseCase:  buyTicketsUseCase,
		createSpotsUseCase: createSpotsUseCase,
		listSpotsUseCase:   listSpotsUseCase,
	}
}

func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")

	input := usecases.GetEventInputDTO{ID: eventID}

	output, err := h.getEventUseCase.Execute(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input usecases.CreateEventInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	var input usecases.BuyTicketsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) CreateSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	var input usecases.CreateSpotsInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input.EventID = eventID

	output, err := h.createSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (h *EventsHandler) writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// CreateSpotsRequest represents the input for creating spots.
type CreateSpotsRequest struct {
	NumberOfSpots int `json:"number_of_spots"`
}

func (h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecases.ListSpotsInputDTO{EventID: eventID}

	output, err := h.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
