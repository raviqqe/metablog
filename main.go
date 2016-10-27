package main

import (
  "log"
  "net/http"
  "os"
  "time"

  "github.com/docopt/docopt-go"
)



func main() {
  args := get_args()
  a, _ := args["<address>"].(string)
  d, _ := args["<document_root>"].(string)
  serve(a, d)
}


func get_args() map[string]interface{} {
  usage := `Usage:
  metablog [-a <address>] [-d <document_root>]

Options:
  -h --help           Show this help
  -a --address        Set an listen address [default: 0.0.0.0:80]
  -d --document-root  Set a document root directory
  --version     Show version
`

  args, e := docopt.Parse(usage, nil, true, "metablog 0.0.0", false)
  if e != nil {
    log.Fatal(e)
  }

  if args["--document-root"] == false {
    d, e := os.Getwd()
    if e != nil {
      log.Fatal(e)
    }

    args["<document_root>"] = d
  }

  return args
}


func serve(a string, d string) {
  s := newServer(a, d)
  log.Fatal(s.ListenAndServe())
}


func newServer(a string, d string) http.Server {
  return http.Server {
    Addr: a,
    Handler: NewHandler(d),
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 5 * time.Second,
  }
}
