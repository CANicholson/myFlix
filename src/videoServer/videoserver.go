package main

import (
	"log"
	"net/http"
)

func videoHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/video/"):]
	title := "video/" + file + ".mp4"
	http.ServeFile(w, r, title)
}

func main() {
	http.HandleFunc("/video/", videoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
