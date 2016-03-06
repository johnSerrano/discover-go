package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Int31() % 100
	if number > 50 {
		fmt.Printf("My number is %d and it is greater than 50\n", number)
	} else {
		fmt.Printf("My number is %d and it is less than or equal to 50\n", number)
	}

	school := "Holberton School"

	if school == "Holberton School" {
		fmt.Println("I am a student at the Holberton School")
	} else {
		//behavior is not defined
	}

	beautifulWeather := true

	if beautifulWeather {
		fmt.Println("It's a beautiful weather!")
	} else {
		//behavior is not defined
	}

	founders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

	for i := 0; i < len(founders); i++ {
		fmt.Println(founders[i], "is a founder")
	}
}
