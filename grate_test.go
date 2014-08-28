package grate

import (
	"fmt"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {

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
