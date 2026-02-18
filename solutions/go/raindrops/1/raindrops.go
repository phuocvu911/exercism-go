package raindrops
import "strconv"
func Convert(number int) string {
	switch {
        case number%105 == 0:
        	return "PlingPlangPlong"
        case number%35 == 0:
        	return "PlangPlong"
        case number%21 ==0:
        	return "PlingPlong"
        case number%15==0:
        	return "PlingPlang"
        case number%7==0:
        	return "Plong"
        case number%5==0:
        	return "Plang"
        case number%3==0:
        	return "Pling"
        default:
        	return strconv.Itoa(number)
    }
}
