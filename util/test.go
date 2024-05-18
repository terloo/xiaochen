package util

import (
	"time"
)

type SpyClock struct {
	SpyTimeStr string
	SpyTime    time.Time
}

func (s *SpyClock) Now() time.Time {
	if s.SpyTime.IsZero() {
		if s.SpyTimeStr == "" {
			s.SpyTime = time.Now()
		} else {
			parse, err := time.Parse(DateLayout, s.SpyTimeStr)
			if err != nil {
				return time.Now()
			}
			s.SpyTime = parse
		}
	}
	if s.SpyTime.IsZero() {
		return time.Now()
	}
	return s.SpyTime
}

func (s *SpyClock) After(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	ch <- s.SpyTime.Add(d)
	return ch
}

var _ Clock = (*SpyClock)(nil)
