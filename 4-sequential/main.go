package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Search for results from omdbapi searches
type Search struct {
	Search []SearchListing `json:"Search"`
}

//SearchListing individual listing from search
type SearchListing struct {
	Title  string `json:"Title"`
	Year   int    `json:"Year,string"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func main() {
	url := "http://www.omdbapi.com/?s=Matrix"
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	var decodedSearch Search
	json.NewDecoder(resp.Body).Decode(&decodedSearch)

	length := len(decodedSearch.Search)

	for i := 0; i < length; i++ {
		fetchMovie("http://www.omdbapi.com/?i=" + decodedSearch.Search[i].IMDBID + "&plot=short&r=json")
	}
}

func fetchMovie(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	var decodedJSON Movie
	json.NewDecoder(resp.Body).Decode(&decodedJSON)

	rating := int(decodedJSON.IMDBRating * 10)

	fmt.Printf(
		"The movie : %s was released in %s - the IMDB rating is %d%% with %s votes.\n",
		decodedJSON.Title, decodedJSON.Released, rating, decodedJSON.IMDBVotes)
}
