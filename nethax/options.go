package nethax

/**
To add support for a below option (such as Verbose()) to a specific options type (such as ScanPortsOption):
1. Nest the option interface type the return value of the option
2. Implement a setter for the option type on the option's struct
**/

// Verbose sets a verbose flag for a compatible procedure
func Verbose() interface {
	ScanPortsOption
} {
	return &verboseOpt{}
}

type verboseOpt struct{}

func (o *verboseOpt) setScanPortsOption(opts *scanPortsOptions) {
	opts.verbose = true
}

// WithGoroutines sets the number of goroutines to use in a concurrent procedure
func WithGoroutines(n int) interface {
	ScanPortsOption
} {
	return &goroutineOpt{nGoroutines: n}
}

type goroutineOpt struct {
	nGoroutines int
}

func (o *goroutineOpt) setScanPortsOption(opts *scanPortsOptions) {
	opts.goroutines = o.nGoroutines
}
