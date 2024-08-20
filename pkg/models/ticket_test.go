package models

import (
	"fmt"
	"testing"
)

func TestGetTicketByMail(t *testing.T) {
	fmt.Println(GetTicketByMail("asdasd@mail.com"))
}
