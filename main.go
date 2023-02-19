package main

import "github.com/uadmin/uadmin"

func main() {
	DBSettings()
	RegisterModels()
	RegisterHandlers()
	ServerConfig()
}

func DBSettings() {
	uadmin.Database = &uadmin.DBSettings{
		Type:     "myslq",
		Name:     "shop",
		User:     "root",
		Password: "Allen is Great 200%",
		Port:     3306,
	}
}

func RegisterModels() {
	uadmin.Register()
}

func RegisterHandlers() {

}

func ServerConfig() {
	uadmin.Port = 8080
	uadmin.RootURL = "/admin/"
	uadmin.StartServer()
}
