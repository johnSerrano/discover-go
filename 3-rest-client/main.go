package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "http://www.omdbapi.com/?i=tt0372784&plot=short&r=json"
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	fmt.Println("status code is", resp.Status)

	var decodedJSON Movie
	json.NewDecoder(resp.Body).Decode(&decodedJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	rating := int(decodedJSON.IMDBRating * 10)

	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes.\n", decodedJSON.Title, decodedJSON.Released, rating, decodedJSON.IMDBVotes)
}
