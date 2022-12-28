package main

import (
	"fmt"

	"net/http"
	"text/template"
)

const port = ":3000"

func Api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact me : ")
}
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./view/" + tmpl + ".page.tmpl")
	if err != err {
		fmt.Println("error")
	}
	t.Execute(w, nil)
}
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact me : ")
}
func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home")
}

func main() {
	http.HandleFunc("/api", Api)
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)

	fmt.Printf("server starting on : http://localhost%s/ !", port)

	http.ListenAndServe(port, nil)
}
