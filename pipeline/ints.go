package pipeline

import "sync"


// Ints returns a channel and sends all values in the given range
// to the channel. The channel is closed when all values have been sent.
func Ints(from, to int) <-chan int {
	out := make(chan int)
	go func() {
		for i := from; i <= to; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// MergeInts merges int output channels into a single channel
func MergeInts(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
