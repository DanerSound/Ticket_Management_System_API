package models

import (
	"fmt"
	"strings"
	"tickets_manager/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Ticket struct {
	gorm.Model
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	TicketNumer int    `json:"numert"`
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&Ticket{})
}

func (ticket *Ticket) CreateTicket() (*Ticket, error) {
	if validityCheck(ticket) {
		db.NewRecord(ticket)
		db.Create(&ticket)
		return ticket, nil
	}
	return nil, fmt.Errorf("invalid ticket")
}

// this function controllo the validity of the input it can be improved
func validityCheck(ticket *Ticket) bool {
	isValid := false
	isValidFullName := len(ticket.Name) >= 2 && len(ticket.Surname) >= 2
	isValidEmail := strings.Contains(ticket.Email, "@")
	if isValidFullName && isValidEmail {
		isValid = true
	}
	return isValid
}

func GetAllTickets() []Ticket {
	var Tickets []Ticket

	db.Find(&Tickets)
	return Tickets
}

func GetTicketByMail(mail string) (*Ticket, *gorm.DB) {
	var ticket Ticket
	db := db.Where("email=?", mail).Find(&ticket)
	if db.RowsAffected == 0 {
		return nil, db
	}
	return &ticket, db
}

func DeleteTicket(email string) Ticket {
	ticketDetails, _ := GetTicketByMail(email)
	db.Where("email=?", email).Delete(ticketDetails)
	return *ticketDetails
}
