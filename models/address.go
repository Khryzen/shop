package models

import "github.com/uadmin/uadmin"

type Address struct {
	uadmin.Model
	Blk        string
	Street     string
	Barangay   Barangay
	BarangayID uint
	City       City
	CityID     uint
	Province   Province
	ProvinceID uint
	Region     Region
	RegionID   uint
	ZipCode    string
}

type Barangay struct {
	uadmin.Model
	Name string `uadmin:"required"`
}

type City struct {
	uadmin.Model
	Code         string `uadmin:"required"`
	Name         string `uadmin:"required;search"`
	RegionCode   string `uadmin:"required;filter"`
	ProvinceCode string `uadmin:"required;filter"`
	Href         string `uadmin:"required"`
}

type Province struct {
	uadmin.Model
	Code       string `uadmin:"required"`
	Name       string `uadmin:"required"`
	RegionCode string `uadmin:"required"`
	Href       string `uadmin:"required"`
}

type Region struct {
	uadmin.Model
	Code string `uadmin:"required"`
	Name string `uadmin:"required"`
	Href string `uadmin:"required"`
}
