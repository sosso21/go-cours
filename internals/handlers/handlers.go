package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact me : ")
}
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	template, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	// Capture any error
	if err != nil {
		log.Fatalln(err)
	}
	// Print out the template to std
	template.Execute(w, nil)

}
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact me : ")
}
func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home")
}
