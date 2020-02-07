package main

import (
	"github.com/stevenctl/gohax/nethax"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	host       = kingpin.Arg("host", "Host to scan ports on.").Required().String()
	goroutines = kingpin.Arg("goroutines", "Host to scan ports on.").Int()
)

func main() {
	kingpin.Parse()
	opts := []nethax.ScanPortsOption{nethax.Verbose()}
	if goroutines != nil {
		opts = append(opts, nethax.WithGoroutines(*goroutines))
	}
	ports := nethax.ScanPorts(*host, opts...)
	for range ports { /* read until channel closes */ }
}
