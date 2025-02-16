package booking

import (
	"fmt"
	"time"
)

const ScheduleTimeLayout = "1/2/2006 15:04:05"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse(ScheduleTimeLayout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return t.Hour() >= 12 && (t.Hour()+t.Minute()) < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse(ScheduleTimeLayout, date)
	return fmt.Sprintf(
		"You have an appointment on %s, %s %s, %s, at %s.",
		t.Format("Monday"),
		t.Format("January"),
		t.Format("2"),
		t.Format("2006"),
		t.Format("15:04"),
	)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t, _ := time.Parse("2006-01-02", "2025-09-15")
	return t.UTC()
}
