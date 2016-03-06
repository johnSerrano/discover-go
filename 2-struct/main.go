package main

import (
	"fmt"
	"math"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func hello(u user) {
	fmt.Println("Hello", u.Name)
}

func age(u user) {
	layout := "January 2, 2006"
	uAge, err := time.Parse(layout, u.DOB)
	if err != nil {
		fmt.Println("asdf")
		fmt.Println(err)
		return
	}
	age := int(math.Floor(time.Since(uAge).Hours() / 8760))
	fmt.Printf("Betty Holberton who was born in %s would be %d years old today\n", u.City, age)
}

func main() {
	u := user{
		"Betty Holberton",
		"March 7, 1917",
		"Philidelphia",
	}

	hello(u)
	age(u)
}
