// gigasecond calculates the date after a gigasecond has passed
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond return the time after a gigasecond passed from the given time t
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1_000_000_000)
}
