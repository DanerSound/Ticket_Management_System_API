package controllers

import (
	"net/http"
	"tickets_manager/pkg/models"
	"tickets_manager/pkg/utils"

	"encoding/json"

	"github.com/gorilla/mux"
)

const availableTickets = 5

var remainingTickets = 5

func CreateTickets(w http.ResponseWriter, r *http.Request) {
	createTicket := &models.Ticket{}
	utils.ParseBody(r, createTicket)
	if remainingTickets == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		errorResponse := map[string]string{"Error": " Sorry but no Tickets Available!"}
		res, _ := json.Marshal(errorResponse)
		w.Write(res)
		return
	}
	ticket, err := createTicket.CreateTicket()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		errorResponse := map[string]string{"Error": " inserted values not valid! Check them!"}
		res, _ := json.Marshal(errorResponse)
		w.Write(res)
		return
	}
	res, _ := json.Marshal(ticket)
	w.WriteHeader(http.StatusOK)
	remainingTickets = remainingTickets - 1
	w.Write(res)

}

func GetTickets(w http.ResponseWriter, r *http.Request) {
	ticketsList := models.GetAllTickets()
	res, _ := json.Marshal(ticketsList)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTicketBymail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Email := vars["email"]

	ticketDetails, _ := models.GetTicketByMail(Email)

	if ticketDetails == nil {
		// Se ticketDetails è nil, significa che il ticket non è stato trovato
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "Ticket non trovato"}
		res, _ := json.Marshal(errorResponse)
		w.Write(res)
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(ticketDetails)
		w.Write(res)
	}

}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	ticket := models.DeleteTicket(email)
	res, err := json.Marshal(ticket)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	if remainingTickets != availableTickets {
		remainingTickets = remainingTickets + 1
	}
	w.Write(res)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	var updateTicket = &models.Ticket{}

	utils.ParseBody(r, updateTicket)
	vars := mux.Vars(r)
	Email := vars["email"]

	ticketDetails, db := models.GetTicketByMail(Email)

	if updateTicket.Name != "" {
		ticketDetails.Name = updateTicket.Name
	}
	if updateTicket.Surname != "" {
		ticketDetails.Surname = updateTicket.Surname
	}
	if updateTicket.Email != "" {
		ticketDetails.Email = updateTicket.Email
	}

	db.Save(&ticketDetails)
	res, _ := json.Marshal(ticketDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
