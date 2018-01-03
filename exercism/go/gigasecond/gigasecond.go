/*
Package gigasecond has a function AddGigasecond that takes someone's birthday
and calculates when someone passes a gigasecond (1000000000) of their lives.
*/
package gigasecond

import "time"

// AddGigasecond takes someone's birthday and returns when they pass a
// gigasecond age.
func AddGigasecond(t time.Time) time.Time {
	// const GigaInt = 1000000000
	// var gigaseconds = time.Duration(GigaInt) * time.Second
	return t.Add(1000000000)
	// return t.Add(gigaseconds)
}
