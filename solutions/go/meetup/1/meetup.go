package meetup

import "time"

type WeekSchedule string

const( //for the test
    Teenth WeekSchedule = "teenth"
    First WeekSchedule = "first"
    Second WeekSchedule = "second"
    Third WeekSchedule = "third"
    Fourth WeekSchedule = "fourth"
    Last WeekSchedule = "last"    
)

var lookup = map[WeekSchedule]int{
    First:1,
	Second:2,
    Third:3,
    Fourth:4,
}

func Day(s WeekSchedule, wd time.Weekday, month time.Month, year int) int {
	switch s{
        case Teenth: //iterates days 13–19 to find the matching weekday
        	for i:=13; i<=19; i++{
                t:= time.Date(year, month, i,0,0,0,0, time.UTC)
                if t.Weekday() == wd {
                    return i
                }
            }
        case Last: //finds the last day of the month, then walks backwards up to 6 days
        	lastday:= time.Date(year, month+1,0,0,0,0,0, time.UTC).Day() //last day of the month is day 0 of next month
        	for i:=lastday; i>=lastday-6; i--{
                t:= time.Date(year, month, i,0,0,0,0, time.UTC)
                if t.Weekday() == wd{
                    return i
                }
            }
        default: //counts forward through the month until the nth occurrence of the weekday
        	nth := lookup[s]
        	count:=0
        	for i:=1; i<=31; i++{
            	t:= time.Date(year, month, i,0,0,0,0, time.UTC)
                if t.Weekday() == wd{
                    count++
                    if count == nth{
                        return i
                    }
                }
        }
    }
    return 0
}
