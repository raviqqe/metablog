package main

import (
  "log"
  "net/http"
  "time"

  "github.com/docopt/docopt-go"
)



func main() {
  args := get_args()
  a, _ := args["<address>"].(string)
  serve(a)
}


func get_args() map[string]interface{} {
  usage := `Usage:
  metablog [-a <address>]

Options:
  -h --help     Show this help
  -a --address  Set an listen address [default: 0.0.0.0:80]
  --version     Show version
`

  args, _  := docopt.Parse(usage, nil, true, "metablog 0.0.0", false)
  return args
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
