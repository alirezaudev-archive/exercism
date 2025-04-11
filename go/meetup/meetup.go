// Package meetup handles meetups.
package meetup

import "time"

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Teenth WeekSchedule = 13
	Last   WeekSchedule = -6
)

func Day(w WeekSchedule, wd time.Weekday, m time.Month, y int) int {
	if w == Last {
		m++
	}
	t := time.Date(y, m, int(w), 12, 0, 0, 0, time.UTC)
	return t.Day() + int(wd-t.Weekday()+7)%7
}
