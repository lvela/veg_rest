package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"veg_rest2/api"
	"veg_rest2/client"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func TestUpdateVegetable(t *testing.T) {

	// given
	client := client.VegetableClient{Host: "http://localhost:8080"}
	todo, _ := client.CreateVegetable("foo", "bar")

	// when
	todo.Status = "doing"
	todo.Title = "baz"
	todo.Description = "bing"
	_, err := client.UpdateVegetable(todo)

	// then
	if err != nil {
		t.Error(err)
	}

	todoResult, _ := client.GetVegetable(todo.Id)

	if !reflect.DeepEqual(todo, todoResult) {
		t.Error("returned todo not updated")
	}

	// cleanup
	_ = client.DeleteVegetable(todo.Id)
}

func TestUpdateNonExistantVegetable(t *testing.T) {

	// given
	client := client.VegetableClient{Host: "http://localhost:8080"}
	todo, _ := client.CreateVegetable("foo", "bar")
	_ = client.DeleteVegetable(todo.Id)

	// when
	todo.Status = "doing"
	todo.Title = "baz"
	todo.Description = "bing"
	_, err := client.UpdateVegetable(todo)

	// then
	if err == nil {
		t.Error(err)
	}

}

func TestUpdateVegetablesStatus(t *testing.T) {

	// given
	client := client.VegetableClient{Host: "http://localhost:8080"}
	todo, _ := client.CreateVegetable("foo", "bar")

	// when
	_, err := client.UpdateVegetableStatus(todo.Id, api.DoingStatus)

	// then
	if err != nil {
		t.Error(err)
	}

	todoResult, _ := client.GetVegetable(todo.Id)

	if todoResult.Status != "doing" {
		t.Error("returned todo status not updated")
	}

	// cleanup
	_ = client.DeleteVegetable(todo.Id)
}

func TestUpdateNotFoundVegetablesStatus(t *testing.T) {

	// given
	client := client.VegetableClient{Host: "http://localhost:8080"}
	id := int32(3)
	// when
	_, err := client.UpdateVegetableStatus(id, api.DoingStatus)

	// then
	if err == nil {
		t.Error(err)
	}
}
