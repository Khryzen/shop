package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mbdeguzman/shopping/models"
	"github.com/uadmin/uadmin"
)

type Pagination_City struct {
	Total    int `json:"total"`
	PerPage  int `json:"perPage"`
	Page     int `json:"page"`
	LastPage int `json:"lastPage"`
}

type City struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	RegionCode   string `json:"region_code"`
	ProvinceCode string `json:"province_code"`
	Href         string `json:"href"`
}

type CityList struct {
	Pagination Pagination_City `json:"pagination"`
	Data       []City          `json:"data"`
}

func fetchCities() {
	url := "https://ph-locations-api.buonzz.com/v1/cities"
	req, err := http.NewRequest("GET", url, nil)
	CheckErr(err)
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	respstr := string(body)
	cityList := CityList{}
	err = json.Unmarshal([]byte(respstr), &cityList)
	CheckErr(err)

	cities := []models.City{}
	uadmin.All(&cities)

	for _, cit := range cityList.Data {
		isExisting := false
		for _, d_cit := range cities {
			if cit.ID == d_cit.Code {
				isExisting = true
			}
		}
		if !isExisting {
			city_val := models.City{}
			city_val.Code = cit.ID
			city_val.Name = cit.Name
			city_val.ProvinceCode = cit.ProvinceCode
			city_val.RegionCode = cit.RegionCode
			city_val.Href = cit.Href
			uadmin.Save(&city_val)
			uadmin.Trail(uadmin.INFO, "New city has been added.")
		}
	}
}

type Pagination_Province struct {
	Total    int `json:"total"`
	PerPage  int `json:"perPage"`
	Page     int `json:"page"`
	LastPage int `json:"lastPage"`
}

type Provinces struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RegionCode string `json:"region_code"`
	Href       string `json:"href"`
}

type ProvincesList struct {
	Pagination Pagination  `json:"pagination"`
	Provinces  []Provinces `json:"data"`
}

func fetchProvinces() {
	url := "https://ph-locations-api.buonzz.com/v1/provinces"
	req, err := http.NewRequest("GET", url, nil)
	CheckErr(err)
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	respstr := string(body)
	provinceList := ProvincesList{}
	err = json.Unmarshal([]byte(respstr), &provinceList)
	CheckErr(err)

	uadmin.Trail(uadmin.DEBUG, "Provinces: %v", provinceList)

	prov := []models.Province{}
	uadmin.All(&prov)

	for _, prv := range provinceList.Provinces {
		isExisting := false

		for _, d_prov := range prov {
			if prv.ID == d_prov.Code {
				isExisting = true
			}
		}
		if !isExisting {
			prov_val := models.Province{}
			prov_val.Code = prv.ID
			prov_val.Name = prv.Name
			prov_val.RegionCode = prv.RegionCode
			prov_val.Href = prv.Href
			uadmin.Save(&prov_val)
			uadmin.Trail(uadmin.INFO, "New province has been saved.")
		}
	}
}
