package main

import (
  "log"
  "net/http"
  "path/filepath"
)



type Handler struct {
  document_root string
}

func NewHandler(d string) Handler {
  d, e := filepath.Abs(d)
  if e != nil {
    log.Fatal(e)
  }

  return Handler {
    document_root: d,
  }
}

func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, world!"))
}
