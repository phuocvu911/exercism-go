package meetup

import "time"

type WeekSchedule int

const( //for the test
    Teenth WeekSchedule = 13
    First WeekSchedule = 1
    Second WeekSchedule = 8
    Third WeekSchedule = 15
    Fourth WeekSchedule = 22
    Last WeekSchedule = -6
)
//calculate anchor date, and add the difference of desire wd 
func Day(s WeekSchedule, wd time.Weekday, month time.Month, year int) int {
    if s == Last{
        month++
    }
    t:= time.Date(year, month, int(s),1,0,0,0, time.UTC)
    return t.Day() + int(wd-t.Weekday() + 7)%7
}
