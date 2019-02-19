package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

type Page struct {
	Title string
	Body  []byte
}

type Movie struct {
	Server int    `bson:"server" json:"server"`
	Name   string `bson:"Name" json:"Name"`
	File   string `bson:"file" json:"file"`
	Thumb  string `bson:"thumb" json:"thumb"`
	Uuid   string `bson:"uuid" json:"uuid"`
}

var templates = template.Must(template.ParseFiles("view.html"))

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

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
	title := r.URL.Path[len("/main/"):]
	p, err := loadPage(title)

	renderTemplate(w, "view", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/video/"):]
	title := "video/" + file + ".mp4"
	http.ServeFile(w, r, title)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are at the index")
}

func main() {
	http.HandleFunc("/main/", mainHandler)
	http.HandleFunc("/video/", videoHandler)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
