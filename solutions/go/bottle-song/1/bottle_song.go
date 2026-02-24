package bottlesong
import (
    "fmt"
    "strings"
)
func Recite(startBottles, takeDown int) []string {
	res:= []string{}
    for i:= startBottles; i>startBottles-takeDown; i--{
        current:= Phrase(i)
        next:= Phrase(i-1)
        res= append(res, fmt.Sprintf("%s hanging on the wall,", current))
        res= append(res, fmt.Sprintf("%s hanging on the wall,", current)) 
        res= append(res, "And if one green bottle should accidentally fall,")
        res= append(res, fmt.Sprintf("There'll be %s hanging on the wall.", strings.ToLower(next)))
        if i>startBottles-takeDown+1{
            res = append(res, "")
        }
    }
    return res
}

func Phrase(n int) string{
    switch n {
        case 0:
        	return "no green bottles"
        case 1:
        	return "One green bottle"
        default:
        	return fmt.Sprintf("%s green bottles", Convert(n))
    }
}

func Convert(n int) string {
	words := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",}

	if n >= 0 && n < len(words) {
		return words[n]
	}
	return ""
}