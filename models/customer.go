package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Customer struct {
	uadmin.Model
	FirstName     string `uadmin:"required"`
	LastName      string `uadmin:"required"`
	Birthdate     time.Time
	Address       Address
	AddressID     uint
	ContactNumber string      `uadmin:"required"`
	EmailAddress  string      `uadmin:"required"`
	User          uadmin.User `uadmin:"required"`
	UserID        uint
}

func (c *Customer) String() string {
	return c.FirstName + " " + c.LastName
}
