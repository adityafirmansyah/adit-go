package main

import (
    "html/template"
    "net/http"
    "path"
)

func init() {
    http.HandleFunc("/about/", about)
}

func about(w http.ResponseWriter, r *http.Request) {
    header := Header {"About", ""}

    fph := path.Join("templates", "header.tpl")
    fpf := path.Join("templates", "footer.tpl")
    page := path.Join("pages", "about.tpl")
    tmpl, err := template.ParseFiles(fph, fpf, page)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.ExecuteTemplate(w, "content", header); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
