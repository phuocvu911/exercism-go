package resistorcolortrio
import (
    "math"
    "strconv"
)
var h = map[string]int{
        "black": 0,
		"brown": 1,
		"red": 2,
		"orange": 3,
		"yellow": 4,
		"green": 5,
		"blue": 6,
		"violet": 7,
		"grey": 8,
		"white": 9,
    }
var mpre = []string{"", "kilo", "mega", "giga"}

func Value(colors []string) int {
    return (h[colors[0]]*10 + h[colors[1]]) * int(math.Pow(10.0,float64(h[colors[2]])))
}

func Label(colors []string) string {
	a:= Value(colors)
    count := 0
    for a>0{
        x:= a%1000
        if x == 0{
            count++
        } else{
            break
        }
        a/=1000
    }
    res := strconv.Itoa(a)
    return res + " " + mpre[count] + "ohms"
}