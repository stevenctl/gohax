package cmd

import (
	"fmt"
	"github.com/stevenctl/gohax/nethax"
)

func PortScan(host string) {
	ports := nethax.ScanPorts(host)
	for p := range ports {
		fmt.Println(p)
	}
}
