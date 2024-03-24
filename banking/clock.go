package banking

import "time"

type clock interface {
	now() time.Time
}

type defaultClock struct{}

func (c *defaultClock) now() time.Time {
	return time.Now()
}
