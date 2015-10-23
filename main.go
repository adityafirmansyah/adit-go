package main

import (
    "html/template"
    "net/http"
    "path"
)

type Header struct {
    Title string
    Description string
}

func init() {
    http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
    header := Header {"Index", "Yamete....!!"}

    fph := path.Join("templates", "header.tpl")
    fpi := path.Join("templates", "index.tpl")
    fpf := path.Join("templates", "footer.tpl")
    tmpl, err := template.ParseFiles(fph, fpi, fpf)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.ExecuteTemplate(w, "content", header); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
