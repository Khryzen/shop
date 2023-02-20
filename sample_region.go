package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mbdeguzman/shopping/models"
	"github.com/uadmin/uadmin"
)

type Regions struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

type Pagination struct {
	Total    int `json:"total"`
	PerPage  int `json:"perPage"`
	Page     int `json:"page"`
	LastPage int `json:"lastPage"`
}

type RegionsResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []Regions  `json:"data"`
}

func fetchRegion() {
	url := "https://ph-locations-api.buonzz.com/v1/regions"
	req, err := http.NewRequest("GET", url, nil)
	CheckErr(err)
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	respstr := string(body)
	var data RegionsResponse
	err = json.Unmarshal([]byte(respstr), &data)
	CheckErr(err)

	regs := []models.Region{}
	uadmin.All(&regs)
	for _, reg := range data.Data {
		isExisting := false
		for _, d_reg := range regs {
			if reg.ID == d_reg.Code {
				isExisting = true
				return
			}
		}
		if !isExisting {
			reg_val := models.Region{}
			reg_val.Code = reg.ID
			reg_val.Name = reg.Name
			reg_val.Href = reg.Href
			uadmin.Save(&reg_val)
			uadmin.Trail(uadmin.INFO, "New region has been added.")
		}
	}
}
