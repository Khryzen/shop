package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mbdeguzman/shopping/models"
	"github.com/uadmin/uadmin"
)

type Provinces struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RegionCode string `json:"region_code"`
	Href       string `json:"href"`
}

type ProvincesList struct {
	Provinces []Provinces `json:"data"`
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
