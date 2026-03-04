package anagram
import "strings"
func Detect(subject string, candidates []string) []string {
	res := []string{}
    s:= Low(subject)
    h:= make(map[rune]int)
    for _,v:= range s{
        h[v]++
    }
    Outer:
    for _, candidate := range candidates{
        can:= Low(candidate)
        p:= make(map[rune]int)
        for _, c:= range can{
            if h[c]==0{
                break
            }
            p[c]++
        }
        for key,val:= range h{
            if p[key] != val{
                continue Outer // need it to skip 2 for loops
            }
        }
        if len(subject) == len(candidate) && len(h) == len(p) && s != can{
            res = append(res, candidate)
        }
    }
    return res
}

func Low(s string) string{
    return strings.ToLower(s)
}
