package models

import "github.com/uadmin/uadmin"

type Shop struct {
	uadmin.Model
	Name          string
	Slogan        string
	Icon          string `uadmin:"image"`
	Logo          string `uadmin:"image"`
	Address       string
	ContactNumber string
	EmailAddress  string `uadmin:"email"`
	Footer        string
	Active        bool
}

func (s Shop) String() string {
	return s.Name
}

func (s Shop) Validate() map[string]interface{} {
	shop := Shop{}
	c := map[string]interface{}{}
	if s.Active {
		if uadmin.Count(&shop, "active = ?", true) > 0 {
			c["Active"] = "Only 1 active shop is allowed."
		}
	}
	return c
}
