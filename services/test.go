package services

import (
  "encoding/json"
  "github.com/danielhood/gosvctest/dataAccess"
)

type Test interface {
  Get() []byte
}

func NewTest() Test {
  return &TestService{}
}

type TestService struct {}

func(* TestService) Get() []byte {
  //return "this is a test"
  tests := dataAccess.ReadAll()

  barr, _ := json.Marshal(tests)
  return barr
}
