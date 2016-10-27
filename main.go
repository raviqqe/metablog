package main

import (
  "net/http"
  "log"
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
    ReadTimeout: 50,
    WriteTimeout: 50,
  }
}
