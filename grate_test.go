package grate

import (
	"fmt"
	"testing"
	"time"
)

func TestTry(_ *testing.T) {

	r := NewRateLimiter(3, time.Second)
	start := time.Now()

	for i := 0; i < 10; i++ {
		if r.Try() {
			fmt.Println("ok!", i, time.Since(start))
		} else {
			fmt.Println("toosoon!", i, time.Since(start))
			time.Sleep(time.Second / 2)
		}
	}

}

func TestWait(_ *testing.T) {

	r := NewRateLimiter(5, time.Millisecond)
	start := time.Now()

	for i := 0; i < 10; i++ {
		r.Wait()
		fmt.Println(i, time.Since(start))
	}

}
