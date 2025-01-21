package main

import (
	"html/template"
	"net/http"
)

var (
	templates = make(map[string]*template.Template)
)

func init() {
	layout := template.Must(template.New("layout").ParseFiles("templates/layout.html"))
	header := template.Must(template.New("header").ParseFiles("templates/partials/header.html"))
	footer := template.Must(template.New("footer").ParseFiles("templates/partials/footer.html"))

	templates["layout"] = layout
	templates["header"] = header
	templates["footer"] = footer
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates[tmpl].Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "layout", map[string]interface{}{
			"Title": "Home",
		})
	})

	http.ListenAndServe(":8080", nil)
}