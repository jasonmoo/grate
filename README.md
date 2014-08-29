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

for i := 0; i < 10; i++ {
	r.Wait()
	fmt.Println(i, time.Since(start))
}
```
```
// Output:
ok! 0 630ns
ok! 1 14.89us
ok! 2 21.232us
toosoon! 3 27.518us
toosoon! 4 501.197154ms
ok! 5 1.001419939s
ok! 6 1.001493119s
ok! 7 1.001513744s
toosoon! 8 1.001533674s
toosoon! 9 1.502730084s
0 2.003078104s
1 2.003151729s
2 2.003167109s
3 3.003233304s
4 3.003304184s
5 3.003320959s
6 4.004411139s
7 4.004498979s
8 4.004520914s
9 5.005519241s
```

####More information
`grate` offers an upper bound on a number of actions per time segment, but not a constant rate.  E.g.  At a rate of 1 per sec, in a 5 second range, no more than 5 actions will take place.  But in any given second up to 2 actions could occur if the first action lands just before the second boundary, and the next action lands just after the second boundary.  Contrast this with a rate limiter that checks the clock on every action and counts how many actions have taken place in the last wall clock second in order to rate limit.  Over time grate will produce a consistent rate throttle.