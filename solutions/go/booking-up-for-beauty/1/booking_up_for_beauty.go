package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
    t, _:= time.Parse(layout,date)
    return t 
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t:= time.Now()
    layout := "January 2, 2006 15:04:05"
    t1, _:= time.Parse(layout,date)
    x:= t.Compare(t1)
    if x == 1 {
        return true
    } else {
    	return false 
}}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    t1, _:= time.Parse(layout,date)
	t:= t1.Hour()
    if t>= 12 && t< 18{
        return true
    }
    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t, _:= time.Parse(layout,date)
    formatTime := t.Format("Monday, January 2, 2006, at 15:04")
    return fmt.Sprintf("You have an appointment on %s.",formatTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15,0,0,0,0, time.UTC)
    
    
    
}
