package handlers

import (
  "net/http"
)

type Ping struct {}

func NewPing() *Ping {
  return &Ping{}
}

func (h *Ping) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "GET":
        s := "pong"
        w.Write([]byte(s))
    default:
        http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    }
}
