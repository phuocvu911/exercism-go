package piglatin
import "strings"
var vow = map[byte]bool{
    'a': true,
    'i': true,
    'e': true,
    'o': true,
    'u': true,
}

func Sentence(s string) string {
	if strings.Contains(s, " "){
        words:= strings.Fields(s)
        res:= ""
        for i, w:= range words{
            res += Word(w)
            if i<len(words)-1{
                res+=" "
            }
        }
        return res
    }
    return Word(s)
}

func Word(s string) string {
	if vow[s[0]] || s[:2] == "xr" || s[:2] == "yt"{
        return s+"ay"
    }
    if !vow[s[0]]{
        for i:= 1; i<len(s); i++{
            if vow[s[i]]{
                if s[i-1] == 'q'&& s[i] == 'u'{
                    return s[i+1:] + s[:i+1] +"ay"
                }
                return s[i:] + s[:i] +"ay"
            }
            if s[i] == 'y'{
                return s[i:] + s[:i] +"ay"
            }
        }
    }
    return s
}
