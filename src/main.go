package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	Server int    `bson:"server" json:"server"`
	Name   string `bson:"Name" json:"Name"`
	File   string `bson:"file" json:"file"`
	Thumb  string `bson:"thumb" json:"thumb"`
	Uuid   string `bson:"uuid" json:"uuid"`
	IP     string
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
		if result[i].Server == 1 {
			result[i].IP = "35.189.90.164"
		} else {
			result[i].IP = "35.230.159.128"
		}
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
	video := r.URL.Path[len("/video/"):]
	session, err := mgo.Dial("mongodb://callum:help@35.246.67.38:27017")
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB("myflix").C("videos")
	var result []Movie
	err = c.Find(bson.M{"uuid": video}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	for i := range result {
		result[i].File = strings.TrimSuffix(result[i].File, filepath.Ext(result[i].File))
		if result[i].Server == 1 {
			result[i].IP = "35.189.90.164"
		} else {
			result[i].IP = "35.230.159.128"
		}
	}
	renderTemplate(w, "video", &result)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/main/", http.StatusFound)
}

func main() {
	http.HandleFunc("/main/", mainHandler)
	http.HandleFunc("/video/", videoHandler)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
