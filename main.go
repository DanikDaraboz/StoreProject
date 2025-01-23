package main

import (
    "html/template"
    "net/http"
    "log"
)

var templates = template.Must(template.ParseFiles(
    "templates/layout.html",
    "templates/partials/header.html",
    "templates/partials/footer.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Println("Template error:", err)
    }
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "layout.html", map[string]interface{}{
            "Title": "Home",
        })
    })

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}