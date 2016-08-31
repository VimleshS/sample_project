package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/save", save)

	log.Println("Listening...")
	http.ListenAndServe(":5000", nil)
}

func save(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	log.Println(name)
	log.Println(r.FormValue("password"))
}
func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("templates", "layout.html")

	var fp string
	if r.URL.Path == "/" {
		fp = path.Join("templates", "/example.html")
	} else {
		fp = path.Join("templates", r.URL.Path)
	}
	log.Printf(" %s | %s", r.URL.Path, fp)

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
