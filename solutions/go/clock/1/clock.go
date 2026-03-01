package clock
import "fmt"
// Define the Clock type here.
type Clock struct{
    Hour int
    Min int
}

func New(h1, m int) Clock { 
    h:= h1%24
    
    for m>=60{
        h++
        m-=60
    }
    
    for m<0{
        h--
        m += 60
    }
    
    for h<0{
        h= 24+h
    }
    
    for h>=24{
        h-=24
    }
    return Clock{h,m}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hour, c.Min + m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hour, c.Min -m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Min)
}
