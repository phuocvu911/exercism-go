package pangram
import "strings"
func IsPangram(input string) bool {
	s:= strings.ToLower(input)
    seen := make(map[rune]bool)
    for _, v:= range s{
        if v >= 'a' && v <= 'z'{
            seen[v] = true
        }     
    }
    return len(seen) == 26
}