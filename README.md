#grate
a rate limiter in 50 lines of go

grate uses a buffered channel to handle rate limiting over time.  This affords concurrent access without mutexing.  

```go
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
```
```
// Output:
ok! 0 615ns
ok! 1 12.314us
ok! 2 15.533us
toosoon! 3 18.16us
toosoon! 4 501.149373ms
ok! 5 1.001422731s
ok! 6 1.001517816s
ok! 7 1.00154813s
toosoon! 8 1.001554808s
toosoon! 9 1.501923701s
```

####More information
`grate` offers an upper bound on a number of actions per time segment, but not a consistent rate.  E.g.  At a rate of 1 per sec, in a 5 second range, no more than 5 actions will take place.  But in any given second up to 2 actions could occur if the first action lands just before the second boundary, and the next action lands just after the second boundary.  Contrast this with a rate limiter that checks the clock on every action and counts how many actions have taken place in the last wall clock second in order to rate limit.  Over time grate will produce a consistent rate throttle.