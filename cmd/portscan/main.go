package main

import (
	"github.com/stevenctl/gohax/nethax"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	host = kingpin.Arg("host", "Host to scan ports on.").Required().String()
)

func main() {
	kingpin.Parse()
	nethax.ScanPorts(*host)
}
