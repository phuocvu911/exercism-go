package pangram
import "strings"
func IsPangram(input string) bool {
	s:= strings.ToLower(input)
    seen := make(map[rune]bool)
    for _, v:= range s{
        if v < 'a' || v > 'z'{
            continue
        }
        seen[v] = true
    }
    if len(seen) != 26{
        return false
    }
    return true
}