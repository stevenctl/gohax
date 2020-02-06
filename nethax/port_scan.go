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


// scan port checks if
func scanPorts(host string, ports <-chan int) <- chan int {
	out := make(chan int)
	go func() {
		for port := range ports {
			fmt.Printf("checking %d\n", port)
			conn, err := net.Dial("tcp", host + ":" + strconv.Itoa(port))
			if err != nil {
				continue
			}
			_ = conn.Close()
			out <- port
		}
		close(out)
	}()
	return out
}

func ScanPorts(host string) <- chan int {
	allPorts := pipeline.Ints(1, maxPort)
	validChans := make([]<- chan int, 1024)
	for i := 0; i < 1024; i ++ {
		validChans[i] = scanPorts(host, allPorts)
	}
	return pipeline.MergeInts(validChans...)
}
