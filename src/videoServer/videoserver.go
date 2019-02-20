package main

import (
	"log"
	"net/http"
)

func thumbHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/thumb/"):]
	title := "thumb/" + file + ".png"
	http.ServeFile(w, r, title)
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/video/"):]
	title := "video/" + file + ".mp4"
	http.ServeFile(w, r, title)
}

func main() {
	http.HandleFunc("/video/", videoHandler)
	http.HandleFunc("/thumb/", thumbHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
