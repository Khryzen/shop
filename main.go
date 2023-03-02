package main

import (
	"net/http"

	"github.com/mbdeguzman/shopping/models"
	"github.com/mbdeguzman/shopping/views"
	"github.com/uadmin/uadmin"
)

func main() {
	DBSettings()
	RegisterModels()
	RegisterHandlers()
	InitialData()
	ServerConfig()
}

func DBSettings() {
	uadmin.Database = &uadmin.DBSettings{
		Type:     "mysql",
		Name:     "shop",
		User:     "root",
		Password: "Allen is Great 200%",
		Port:     3306,
	}
}

func RegisterModels() {
	uadmin.Register(
		models.Address{},
		models.Barangay{},
		models.City{},
		models.Category{},
		models.Customer{},
		models.Item{},
		models.ItemImage{},
		models.Shop{},
		models.Province{},
		models.Region{},
	)
}

func RegisterHandlers() {
	http.HandleFunc("/", uadmin.Handler(views.IndexHandler))
}

func ServerConfig() {
	uadmin.Port = 8080
	uadmin.RootURL = "/admin/"
	uadmin.StartServer()
}
