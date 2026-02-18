package bob
import "strings"

// Hey should have a comment documenting it.
func Hey(in string) string {
	s:= strings.TrimSpace(in)
    switch{
    	case s == "":
        	return "Fine. Be that way!" 
        case s[len(s)-1] == '?' && IsCapital(s) && !Noletter(s):
        	return "Calm down, I know what I'm doing!" 
        case s[len(s)-1] == '?':
        	return "Sure."
        
        case IsCapital(s) && !Noletter(s):
        	return "Whoa, chill out!" 
        default:
        	return "Whatever."
    }
}

func IsCapital(s string) bool{
    if strings.ToUpper(s) == s {
        return true
    }
    return false
}

func Noletter(s string) bool {
    for _,v:= range s{
        if v>='a' && v<='z' || v>='A' && v<='Z'{
            return false
        }
    }
    return true
}
