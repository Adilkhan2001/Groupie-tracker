package internal

import (
	"fmt"
	"log"
	"net/http"
)

func Starter() {
	host := ":8080"
	mux := http.NewServeMux()
	styles := http.FileServer(http.Dir("./templates/static/"))

	handler := http.StripPrefix("/static/", styles)
	mux.Handle("/static/", handler)

	mux.HandleFunc("/", MainPageHandler)
	mux.HandleFunc("/artist/", ViewArtistHandler)
	fmt.Println("Listenning: http://localhost:8080/")
	err := http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatal("ListenAndServer: ", err, "...")
	}
}
