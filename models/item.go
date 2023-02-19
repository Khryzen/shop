package models

import "github.com/uadmin/uadmin"

type Item struct {
	uadmin.Model
	Name        string   `uadmin:"search;required"`
	Category    Category `uadmin:"filter;required"`
	CategoryID  uint
	Description string  `uadmin:"required"`
	Price       float64 `uadmin:"required"`
	Featured    bool    `uadmin:"search;required"`
	Hidden      bool    `uadmin:"search;required"`
	Quantity    int     `uadmin:"edit:read_only"`
	Remaining   int     `uadmin:"read_only"`
}

func (i *Item) String() string {
	return i.Name
}
