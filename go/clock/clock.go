package clock

import (
	"fmt"
)

type Clock struct {
	hours, minutes int
}

func divmod(numerator, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator
}

func mod(a, m int) int {
	return ((a % m) + m) % m
}

func normalizeTime(h, m int) (int, int) {
	mh, m := divmod(m, 60)
	h += mh
	if m < 0 { // abs(m) less than 60 after mod
		m += 60
		h--
	}
	h = mod(h, 24)
	return h, m
}

func New(h, m int) Clock {
	h, m = normalizeTime(h, m)
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	c.hours, c.minutes = normalizeTime(c.hours, c.minutes+m)
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.hours, c.minutes = normalizeTime(c.hours, c.minutes-m)
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
