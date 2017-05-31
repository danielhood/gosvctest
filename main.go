// Ref: https://github.com/komand/gosea

package main

import (
  "log"
  "net/http"
  "github.com/danielhood/gosvctest/handlers"
)

func main () {
  log.Print("wazzup!")

  // openssl genrsa -out server.key 2048
  certPath := "server.pem"

  // openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
  keyPath := "server.key"

  // http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	//    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  // })

  pingHandler := handlers.NewPing()
  testHandler := handlers.NewTest()

  http.Handle("/ping", pingHandler)
  http.Handle("/test", testHandler)

  //log.Fatal(http.ListenAndServe(":8080", nil))
  log.Fatal(http.ListenAndServeTLS(":8080", certPath, keyPath, nil))
}
