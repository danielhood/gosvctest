package handlers

import (
  "net/http"

  "github.com/danielhood/gosvctest/services"
)

type Test struct {
  svc services.Test
}

func NewTest() *Test {
  return &Test{services.NewTest()}
}

func (h *Test) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "GET":
        barr := h.svc.Get()
        w.Write(barr)
    default:
        http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    }
}
