package models

import "github.com/uadmin/uadmin"

type Category struct {
	uadmin.Model
	Name     string `uadmin:"required"`
	Featured bool
	Active   bool
	Hidden   bool
}

func (c *Category) String() string {
	return c.Name
}

func (c *Category) Validate() map[string]interface{} {
	category := Category{}
	errMsg := map[string]interface{}{}
	if c.Featured {
		if uadmin.Count(&category, "featured = ?", true) > 0 {
			errMsg["Featured"] = "Only 1 category is allowed to be featured."
		}
	}
	return errMsg
}
