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
	// uadmin.Trail(uadmin.DEBUG, "CITIES: %v", cityList.Data)

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
