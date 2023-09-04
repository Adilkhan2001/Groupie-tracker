package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Band struct {
	ID              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	Locations       string   `json:"locations"`
	ConcertDatesAPI string   `json:"concertDates"`
	Relations       string   `json:"relations"`
	ConcertDates    []string
}

type (
	Dates    map[string][]string
	Concerts struct {
		ID             int   `json:"id"`
		DatesLocations Dates `json:"datesLocations"`
	}
	Indexx struct {
		Concerts []Concerts `json:"index"`
	}
)

func GetJson() []Band {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalf("Response Json error: %v", err)
	}

	var Bands []Band
	err_json := json.NewDecoder(response.Body).Decode(&Bands)
	if err_json != nil {
		log.Fatalf("Json Decode error: %v", err_json)
	}
	return Bands
}

func GetDates() Indexx {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatalf("GetDates() err")
	}
	var Indexxx Indexx
	err_json := json.NewDecoder(response.Body).Decode(&Indexxx)
	if err_json != nil {
		log.Fatalf("GetDates() err")
	}
	return Indexxx
}
