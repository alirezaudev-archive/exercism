package clock

import (
	"strconv"
)

type Clock struct {
	H int
	M int
}

func New(h, m int) Clock {
	return calcClock(Clock{H: h, M: m})
}

func (c Clock) Add(m int) Clock {
	c.M += m
	return calcClock(c)
}

func (c Clock) Subtract(m int) Clock {
	c.M -= m
	return calcClock(c)
}

func (c Clock) String() string {
	h := strconv.Itoa(c.H)
	m := strconv.Itoa(c.M)

	if len(h) == 1 {
		h = "0" + h
	}

	if len(m) == 1 {
		m = "0" + m
	}

	return h + ":" + m
}

func calcClock(c Clock) Clock {
	minuteResult := calcMinute(c.M)
	c.M = minuteResult.M
	c.H = calcHour(calcHour(c.H).H + calcHour(minuteResult.H).H).H
	return c
}

func calcHour(h int) Clock {
	h = h % 24
	if h < 0 {
		h = 24 + h
	}
	return Clock{H: h}
}

func calcMinute(m int) Clock {
	h := 0
	if m > 59 || m < 0 {
		h = m / 60
		m = m % 60
		if m < 0 {
			h--
			m = 60 + m
		}
	}

	return Clock{H: h, M: m}
}
