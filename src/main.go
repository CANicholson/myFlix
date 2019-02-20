package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"gopkg.in/mgo.v2"
)

type Movie struct {
	Server int    `bson:"server" json:"server"`
	Name   string `bson:"Name" json:"Name"`
	File   string `bson:"file" json:"file"`
	Thumb  string `bson:"thumb" json:"thumb"`
	Uuid   string `bson:"uuid" json:"uuid"`
}

var templates = template.Must(template.ParseFiles("main.html", "video.html"))

func mainHandler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("mongodb://callum:help@35.246.67.38:27017")
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB("myflix").C("videos")
	var result []Movie
	err = c.Find(nil).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	for i := range result {
		result[i].Thumb = strings.TrimSuffix(result[i].Thumb, filepath.Ext(result[i].Thumb))
		result[i].File = strings.TrimSuffix(result[i].File, filepath.Ext(result[i].File))
	}
	renderTemplate(w, "main", &result)
}

func renderTemplate(w http.ResponseWriter, tmpl string, c *[]Movie) {
	err := templates.ExecuteTemplate(w, tmpl+".html", c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "video.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are at the index")
}

func main() {
	http.HandleFunc("/main/", mainHandler)
	http.HandleFunc("/video", videoHandler)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
