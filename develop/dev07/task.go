package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	single := make(chan interface{})
	done := make(chan struct{})
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			select {
			case <-ch:
				done <- struct{}{}
			case <-done:
				return
			}
		}(ch)
	}
	go func() {
		<-done
		close(done)
		close(single)
	}()
	return single
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(20200*time.Millisecond),
		sig(5*time.Minute),
		sig(15*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v\n", time.Since(start))
}
