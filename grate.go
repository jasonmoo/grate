package grate

import "time"

type RateLimiter chan struct{}

func NewRateLimiter(n int, d time.Duration) RateLimiter {
	r := RateLimiter(make(chan struct{}, n))
	go func() {
		for {
		SLEEP:
			time.Sleep(d)
			for i := 0; i < n; i++ {
				select {
				case _, ok := <-r:
					if !ok {
						return
					}
				default:
					goto SLEEP
				}
			}
		}
	}()
	return r
}

func (r RateLimiter) Try() bool {
	select {
	case r <- struct{}{}:
		return true
	default:
		return false
	}
}

func (r RateLimiter) Wait() {
	r <- struct{}{}
}

func (r RateLimiter) Close() {
	close(r)
}
