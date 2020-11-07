package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Breed    string
	WeightKg int `json:",omitempty"`
	// In cases where an empty value does not exist, omitempty is of no use.
	// An embedded struct, for example, does not have an empty value.
	// To solve this, use a struct pointer instead.
	Size *dimension `json:",omitempty"`
}

type dimension struct {
	Height int
	Width  int
}

// The case where you want to differentiate between a default value and a zero value
// We do not want 0 to be empty value for NumberOfCustomers
// One way to solve this is to use an int pointer
type Restaurant struct {
	NumberOfCustomers *int `json:",omitempty"`
}

func main() {
	d := Dog{
		Breed:    "dalmation",
		WeightKg: 45,
	}

	b, _ := json.Marshal(d)
	fmt.Println(string(b))

	d = Dog{
		Breed: "pug",
	}

	b, _ = json.Marshal(d)
	fmt.Println(string(b))

	n := 0
	r := Restaurant{NumberOfCustomers: &n}

	b, _ = json.Marshal(r)
	fmt.Println(string(b))

	r = Restaurant{}

	b, _ = json.Marshal(r)
	fmt.Println(string(b))
}
