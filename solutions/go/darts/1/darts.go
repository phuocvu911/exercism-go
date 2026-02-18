package darts
import "math"

func Score(x, y float64) int {
	coor:= math.Sqrt(x*x + y*y)
    switch {
        case coor <= 1:
        	return 10
        case coor <= 5:
        	return 5
        case coor <= 10:
        	return 1
        default:
        	return 0
    }
}
