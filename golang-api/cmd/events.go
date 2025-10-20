package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gabrielgodoi987/CodeTicket/golang-api/internal/events/infra/repository"
	"github.com/gabrielgodoi987/CodeTicket/golang-api/internal/events/infra/service"
	usecases "github.com/gabrielgodoi987/CodeTicket/golang-api/internal/events/useCases"
	httpHandler "github.com/gabrielgodoi987/CodeTicket/golang-api/internal/events/infra/http" 
)


func main(){
	db, err := sql.Open("mysql", "test_user:test_password@tcp(golang-mysql:3306)/test_db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() 

	eventRepo, err := repository.NewMysqlEventRepository(db)

	if err != nil {
		panic(err)
	}

	// setar as urls dos partners
	partnerBaseURLs := map[int]string{
		1: "http://localhost:9080/partner1",
		2: "http://localhost:9080/partner2",
	}

	// criação dos usecases
	listEventsUseCase := usecases.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecases.NewGetEventUseCase(eventRepo)
	createEventUseCase := usecases.NewCreateEventUseCase(eventRepo)
	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)
	buyTicketsUseCase := usecases.NewBuyTicketsUseCase(eventRepo, partnerFactory)
	createSpotsUseCase := usecases.NewCreateSpotsUseCase(eventRepo)
	listSpotsUseCase := usecases.NewListSpotsUseCase(eventRepo)

	// Handlers HTTP
	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		getEventUseCase,
		createEventUseCase,
		buyTicketsUseCase,
		createSpotsUseCase,
		listSpotsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /events", eventsHandler.CreateEvent)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)
	r.HandleFunc("POST /events/{eventID}/spots", eventsHandler.CreateSpots)
}