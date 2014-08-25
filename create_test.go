package main

import (
	"fmt"
	"log"
	"testing"

	"veg_rest2/client"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func TestCreateVegetable(t *testing.T) {

	// given
	client := client.VegetableClient{Host: "http://localhost:8080"}

	// when
	vegetable, err := client.CreateVegetable("foo", "bar")

	//then
	if err != nil {
		t.Error(err)
	}

	if vegetable.Title != "foo" && vegetable.Description != "bar" {
		t.Error("returned vegetable not right")
	}

	// cleanup
	_ = client.DeleteVegetable(vegetable.Id)
}
