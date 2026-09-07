package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t,_:= time.Parse("1/2/2006 15:04:05",date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout:="January 2, 2006 15:04:05"
    target,_:=time.Parse(layout,date)
    return target.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout:="Monday, January 2, 2006 15:04:05"
    t,_:=time.Parse(layout,date)
    x:=t.Hour()
    if x>=12 && x<18{
        return true
    }else{
        return false
    }
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t:=Schedule(date)
    formatted:=t.Format("Monday, January 2, 2006, at 15:04")
    return fmt.Sprintf("You have an appointment on %s.",formatted)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
