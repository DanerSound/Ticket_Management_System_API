package routes

import (
	"tickets_manager/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterTicketsRoutes = func(router *mux.Router) {
	router.HandleFunc("/ticket/", controllers.CreateTickets).Methods("POST")
	router.HandleFunc("/tickets/", controllers.GetTickets).Methods("GET")
	router.HandleFunc("/ticket/{email}", controllers.GetTicketBymail).Methods("GET")
	router.HandleFunc("/tickets/{email}", controllers.UpdateTicket).Methods("PUT")
	router.HandleFunc("/tickets/{email}", controllers.DeleteTicket).Methods("DELETE")
}

/*
	routes, app, config, main, models, bookcontrolloer
*/
