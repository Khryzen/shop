package models

import "github.com/uadmin/uadmin"

type ItemImage struct {
	uadmin.Model
	Image  string `uadmin:"image"`
	Item   Item   `uadmin:"required"`
	ItemID uint
	Active bool `uadmin:"required;default_value:true"`
}

func (i *ItemImage) String() string {
	uadmin.Preload(i)
	return i.Item.Name
}
