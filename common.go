package main

import "github.com/uadmin/uadmin"

func CheckErr(err error) {
	if err != nil {
		uadmin.Trail(uadmin.ERROR, err)
		return
	}
}
