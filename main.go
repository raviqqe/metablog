package main

import (
  "net/http"
  "log"
)



func main() {
  s := newServer(":80");

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
