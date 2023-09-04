package internal

import (
	"groupie-tracker/api"
	"strconv"
	"strings"
)

func Search(s string) ([]api.Band, []string) {
	Bands := api.GetJson()
	var result []api.Band
	var artists []string
	creation_date_int, err := strconv.Atoi(s)
	if err != nil {
		creation_date_int = -1
	}
	// members    name   creation_date  first album date   locations
	for _, band := range Bands {
		for _, names := range band.Members {
			if strings.Contains(strings.ToLower(names), strings.ToLower(s)) && !ContainStruct(result, band) {
				result = append(result, band)
				artists = append(artists, names)
			}
		}
		if strings.Contains(strings.ToLower(band.FirstAlbum), strings.ToLower(s)) && !ContainStruct(result, band) {

			result = append(result, band)
			artists = append(artists, band.FirstAlbum)
		}
		if strings.Contains(strings.ToLower(band.Name), strings.ToLower(s)) && !ContainStruct(result, band) {

			result = append(result, band)
			artists = append(artists, band.Name)
		}
		if band.CreationDate == creation_date_int && !ContainStruct(result, band) {
			result = append(result, band)
			creation_date_str := strconv.Itoa(band.CreationDate)
			artists = append(artists, creation_date_str)
		}
	}

	// search locations
	Concerts := api.GetDates().Concerts
	id_max := len(Concerts)
	for id := 0; id < id_max; id++ {
		for key, dates := range Concerts[id].DatesLocations {
			println("Search: ", key, " ", s)
			if strings.Contains(strings.ToLower(key), strings.ToLower(s)) && !ContainStruct(result, Bands[id]) {
				result = append(result, Bands[id])
				artists = append(artists, key)
			}
			for _, date := range dates {
				if strings.Contains(strings.ToLower(date), strings.ToLower(s)) && !ContainStruct(result, Bands[id]) {
					result = append(result, Bands[id])
					artists = append(artists, date)
				}
			}
		}
	}

	return result, artists
}

func ContainStruct(Bands []api.Band, Band api.Band) bool {
	for _, band := range Bands {
		if band.ID == Band.ID {
			return true
		}
	}
	return false
}
