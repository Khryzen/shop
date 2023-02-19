package models

import "github.com/uadmin/uadmin"

type Item struct {
	uadmin.Model
	Name        string   `uadmin:"search;required"`
	Category    Category `uadmin:"filter;required"`
	CategoryID  uint
	Description string  `uadmin:"required"`
	Price       float64 `uadmin:"required"`
	Featured    bool    `uadmin:"search"`
	Hidden      bool    `uadmin:"search"`
	Quantity    int     `uadmin:"required;read_only:edit"`
	Remaining   int     `uadmin:"read_only"`
}

func (i *Item) String() string {
	return i.Name
}
