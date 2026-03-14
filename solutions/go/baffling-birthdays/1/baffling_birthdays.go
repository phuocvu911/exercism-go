package bafflingbirthdays

import (
    "time"
    "math/rand"
)

var days float64 = 365.0
func SharedBirthday(dates []time.Time) bool {
	for i,d:= range dates{
        for j:=i+1; j<len(dates); j++{
            if d.Day() == dates[j].Day() && d.Month() == dates[j].Month(){
                return true
            }
        }
    }
    return false
}

func RandomBirthdates(size int) []time.Time {
	res:= []time.Time{}
    min:= time.Date(1970,1,1,0,0,0,0,time.UTC).Unix()
    max:= time.Date(2026,3,14,0,0,0,0,time.UTC).Unix()
    delta:= max-min
    for i:=0; i<size; i++{
        sec:= rand.Int63n(delta) + min
        year := time.Unix(sec,0).Year() //filter out leap year
        if year%4==0{
            i--
            continue
        }
        res = append(res, time.Unix(sec,0))
    }
    return res
}

func EstimatedProbability(size int) float64 {
    if size == 1{
        return 0
    }
    notmatch:= 1.0
    for i:=2;i<=size; i++{
        notmatch*= (days-float64(i) + 1.0)/days
    }
    return (1.0 - notmatch)*100.0  
}
