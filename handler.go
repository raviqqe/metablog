package main

import "net/http"



type Handler struct {}

func NewHandler() Handler {
  return Handler {}
}

func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, world!"))
}
