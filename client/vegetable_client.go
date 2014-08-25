package client

import (
	"log"
	"strconv"

	"veg_rest2/api"
)

var _ = log.Print

type VegetableClient struct {
	Host string
}

func (tc *VegetableClient) CreateVegetable(name string, amount_per_100 string) (api.Vegetable, error) {
	var respVegetable api.Vegetable
	vegetable := api.Vegetable{Name: name, Amount_per_100: amount_per_100}

	url := tc.Host + "/vegetable"
	r, err := makeRequest("POST", url, vegetable)
	if err != nil {
		return respVegetable, err
	}
	err = processResponseEntity(r, &respVegetable, 201)
	return respVegetable, err
}

func (tc *VegetableClient) GetAllVegetables() ([]api.Vegetable, error) {
	var respVegetables []api.Vegetable

	url := tc.Host + "/vegetable"
	r, err := makeRequest("GET", url, nil)
	if err != nil {
		return respVegetables, err
	}
	err = processResponseEntity(r, &respVegetables, 200)
	return respVegetables, err
}

func (tc *VegetableClient) GetVegetable(id int32) (api.Vegetable, error) {
	var respVegetable api.Vegetable

	url := tc.Host + "/vegetable/" + strconv.FormatInt(int64(id), 10)
	r, err := makeRequest("GET", url, nil)
	if err != nil {
		return respVegetable, err
	}
	err = processResponseEntity(r, &respVegetable, 200)
	return respVegetable, err
}

func (tc *VegetableClient) UpdateVegetable(vegetable api.Vegetable) (api.Vegetable, error) {
	var respVegetable api.Vegetable

	url := tc.Host + "/vegetable/" + strconv.FormatInt(int64(vegetable.Id), 10)
	r, err := makeRequest("PUT", url, vegetable)
	if err != nil {
		return respVegetable, err
	}
	err = processResponseEntity(r, &respVegetable, 200)
	return respVegetable, err
}

func (tc *VegetableClient) UpdateVegetableStatus(id int32, status string) (api.Vegetable, error) {
	var respVegetable api.Vegetable

	patchArr := make([]api.Patch, 1)
	patchArr[0] = api.Patch{Op: "replace", Path: "/status", Value: string(status)}

	url := tc.Host + "/vegetable/" + strconv.FormatInt(int64(id), 10)
	r, err := makeRequest("PATCH", url, patchArr)
	if err != nil {
		return respVegetable, err
	}
	err = processResponseEntity(r, &respVegetable, 200)
	return respVegetable, err
}

func (tc *VegetableClient) DeleteVegetable(id int32) error {
	url := tc.Host + "/vegetable/" + strconv.FormatInt(int64(id), 10)
	r, err := makeRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	return processResponse(r, 204)
}
