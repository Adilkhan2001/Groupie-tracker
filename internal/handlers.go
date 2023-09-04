package internal

import (
	"groupie-tracker/api"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type errorResponse struct {
	Code    int
	Name    string
	Message string
}

var errorsResponses = map[int]errorResponse{
	500: {
		500,
		"Internal",
		"",
	},
	404: {
		404,
		"Not found",
		"",
	},
	400: {
		400,
		"Bad Request",
		"",
	},
	405: {
		405,
		"Method Not Allowed",
		"",
	},
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {

		errorHandler(w, r, errorsResponses[404])
		return
	}

	if r.Method == "GET" {
		Bands := api.GetJson()
		list := []string {
			"freddy merc",
			"lamar",
		}
		answer := map[string]interface{}{
			"Band" : Bands,
			"List" : list,
		}
		err := tpl.ExecuteTemplate(w, "main.html", answer)
		if err != nil {
			errorHandler(w, r, errorsResponses[500])
			return
		}
	} else if r.Method == "POST" {
		string_to_search := r.PostFormValue("search")
		if len(string_to_search) == 0 {
			Bands := api.GetJson()
			err := tpl.ExecuteTemplate(w, "main.html", Bands)
			if err != nil {
				errorHandler(w, r, errorsResponses[500])
				return
			}
			return
		}
		Bands, list := Search(string_to_search)
		println(string_to_search)
		answer := map[string]interface{} {
			"Band" : Bands,
			"List" : list,
		}
		err := tpl.ExecuteTemplate(w, "main.html", answer)
		if err != nil {
			errorHandler(w, r, errorsResponses[500])
			return
		}
		// errorHandler(w, r, errorsResponses[405])
		// return
	}
}

func ViewArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, r, errorsResponses[405])
		return
	}
	id_num := r.URL.Path[8:]
	if ContainsZeroes(id_num) || len(id_num) == 0 {
		errorHandler(w, r, errorsResponses[400])
		return
	}
	id, err := strconv.Atoi(strings.Split(r.URL.Path, "/artist/")[1])
	if err != nil {
		errorHandler(w, r, errorsResponses[400])
		return
	}
	Bands := api.GetJson()
	Locations := api.GetDates()

	if id > len(Bands) || id <= 0 {
		errorHandler(w, r, errorsResponses[404])
		return
	}
	ans := map[string]interface{}{
		"Name":      strings.ReplaceAll(strings.ToLower(Bands[id-1].Name), " ", ""),
		"Band":      Bands[id-1],
		"Locations": Locations.Concerts[id-1].DatesLocations,
	}
	err_artist := tpl.ExecuteTemplate(w, "artist.html", ans)
	if err_artist != nil {
		errorHandler(w, r, errorsResponses[500])
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, err errorResponse) {
	w.WriteHeader(err.Code)
	err_internal := tpl.ExecuteTemplate(w, "error.html", err)
	if err_internal != nil {
		println("true")
		http.Error(w, "Internal server error", 500)
	}
}

// func redirectToHome(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
// }
