package main

import (
    "html/template"
    "net/http"
    "path"
)

func init() {
    http.HandleFunc("/avatar", avatar)
}

func avatar(w http.ResponseWriter, r *http.Request) {
    header := Header {"Index", ""}

    fph := path.Join("templates", "header.tpl")
    fpf := path.Join("templates", "footer.tpl")
    page := path.Join("pages", "avatar.tpl")
    tmpl, err := template.ParseFiles(fph, fpf, page)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.ExecuteTemplate(w, "content", header); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
