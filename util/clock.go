package util

import "time"

type Clock interface {
	Now() time.Time
	After(d time.Duration) <-chan time.Time
}

type RealClock struct {
}

var _ Clock = (*RealClock)(nil)

func (*RealClock) Now() time.Time {
	return time.Now()
}

func (*RealClock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}
