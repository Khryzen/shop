package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mbdeguzman/shopping/models"
	"github.com/uadmin/uadmin"
)

type Barangay struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	RegionCode   string `json:"region_code"`
	ProvinceCode string `json:"province_code"`
	CityCode     string `json:"city_code"`
	Href         string `json:"href"`
}

type BarangayResponse struct {
	Data []Barangay `json:"data"`
}

func fetchBarangays() {
	url := "https://ph-locations-api.buonzz.com/v1/barangays"
	req, err := http.NewRequest("GET", url, nil)
	CheckErr(err)
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	respstr := string(body)
	var data BarangayResponse
	err = json.Unmarshal([]byte(respstr), &data)
	CheckErr(err)

	bars := []models.Barangay{}
	uadmin.All(&bars)
	for _, barangay := range data.Data {
		isExisting := false
		for _, d_barangay := range bars {
			if barangay.ID == d_barangay.Code {
				isExisting = true
				return
			}
		}
		if !isExisting {
			bar_val := models.Barangay{}
			bar_val.Code = barangay.ID
			bar_val.Name = barangay.Name
			bar_val.RegionCode = barangay.RegionCode
			bar_val.ProvinceCode = barangay.ProvinceCode
			bar_val.CityCode = barangay.CityCode
			bar_val.Href = barangay.Href
			uadmin.Save(&bar_val)
			uadmin.Trail(uadmin.INFO, "New barangay has been added.")
		}
	}
}
