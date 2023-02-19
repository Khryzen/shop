package models

import "github.com/uadmin/uadmin"

type ItemImage struct {
	uadmin.Model
	Image  string `uadmin:"image"`
	Item   Item   `uadmin:"required"`
	ItemID uint
	Active bool
}

func (i *ItemImage) String() string {
	uadmin.Preload(i)
	return i.Item.Name
}
