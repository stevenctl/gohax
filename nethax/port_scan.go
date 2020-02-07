package nethax

import (
	"fmt"
	"github.com/stevenctl/gohax/internal/pipeline"
	"net"
	"strconv"
)

const (
	maxPort = 65535
)

// ScanPortsOption allows option funcs to set scanPortsOptions values
type ScanPortsOption interface {
	setScanPortsOption(options *scanPortsOptions)
}

type scanPortsOptions struct {
	goroutines int
	verbose    bool
}

func defaultScanPortsOptions() *scanPortsOptions {
	return &scanPortsOptions{
		goroutines: 128,
		verbose:    false,
	}
}

// ScanPorts will find all open ports on a host
func ScanPorts(host string, options ...ScanPortsOption) <-chan int {
	opts := defaultScanPortsOptions()
	for _, opt := range options {
		opt.setScanPortsOption(opts)
	}

	allPorts := pipeline.Ints(1, maxPort)
	foundPorts := make([]<-chan int, 128)
	for i := 0; i < 128; i++ {
		foundPorts[i] = scanPorts(host, allPorts, opts.EagerPrint)
	}
	return pipeline.MergeInts(foundPorts...)
}

func scanPorts(host string, ports <-chan int, print bool) <-chan int {
	out := make(chan int)
	go func() {
		for port := range ports {
			conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
			if err != nil {
				continue
			}
			_ = conn.Close()
			if print {
				fmt.Println(port)
			}
			out <- port
		}
		close(out)
	}()
	return out
}
