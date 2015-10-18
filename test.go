package main

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/avatar", avatar)
}

func root(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, indexHtml)
}

const indexHtml = `
<html>
  <body>
    <h1>Yamete....!!</h1>
  </body>
</html>
`

func avatar(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, avatarHtml)
}

const avatarHtml = `
<html>
  <body>
    <img src="images/avatar.jpg" alt="">
  </body>
</html>
`
