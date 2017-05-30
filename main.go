package main

import (
  "log"
  "net/http"
  "github.com/danielhood/gosvctest/handlers"
)

func main () {
  log.Print("wazzup!")

  certPath := "server.pem"
  keyPath := "server.key"

  // http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	//    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  // })

  handler := NewPing()
  http.Handle("/ping", handler)

  //log.Fatal(http.ListenAndServe(":8080", nil))
  log.Fatal(http.ListenAndServeTLS(":8080", certPath, keyPath, nil))
}
