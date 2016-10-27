package main

import (
  "log"
  "net/http"
  "path/filepath"
)



type Handler struct {
  document_root, command string
}

func NewHandler(d, c string) Handler {
  d, e := filepath.Abs(d)
  if e != nil {
    log.Fatal(e)
  }

  return Handler {
    document_root: d,
    command: c,
  }
}

func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, world!"))
}
