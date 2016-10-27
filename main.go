package main

import (
  "log"
  "net/http"
  "time"
)



func main() {
  serve(":80")
}


func serve(a string) {
  s := newServer(a)
  log.Fatal(s.ListenAndServe())
}


func newServer(a string) http.Server {
  return http.Server {
    Addr: a,
    Handler: NewHandler(),
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 5 * time.Second,
  }
}
