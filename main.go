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
  c, _ := args["<command>"].(string)
  serve(a, d, c)
}


func get_args() map[string]interface{} {
  usage := `Usage:
  metablog [-a <address>] [-d <directory>] <command>

Options:
  -h, --help  Show this help.
  -a <address>, --address <address>
              Set an listen address. [default: 0.0.0.0:80]
  -d <directory>, --document-root <directory>
              Set a document root directory. The default is a current working
              directory.
  --version   Show version.
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


func serve(a, d, c string) {
  s := newServer(a, d, c)
  log.Fatal(s.ListenAndServe())
}


func newServer(a, d, c string) http.Server {
  return http.Server {
    Addr: a,
    Handler: NewHandler(d, c),
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 5 * time.Second,
  }
}
