package booking

import (
	"fmt"
	"time"
)

func ParseTime(layout, date string) time.Time {
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return t
}

// Schedule returns a time.Time from a string containing a date.
// "7/25/2019 13:45:00"
func Schedule(date string) time.Time {
	return ParseTime("1/2/2006 15:04:05", date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return ParseTime("January 2, 2006 15:04:05", date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	hour := ParseTime("Monday, January 2, 2006 15:04:05", date).Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	formattedDate := Schedule(date).Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", formattedDate)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
