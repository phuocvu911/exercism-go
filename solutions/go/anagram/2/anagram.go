package anagram
import "strings"
import "sort"
func Detect(subject string, candidates []string) []string {
	res := []string{}
    for _, can := range candidates{
        if IsAnagram(subject, can){
            res = append(res, can)
        }
    }
    return res
}

func IsAnagram(a, b string) bool{
    return strings.ToLower(a) != strings.ToLower(b) && SortStr(a) == SortStr(b)
}

func SortStr(s string) string{
    a:= strings.Split(strings.ToLower(s), "")
    sort.Strings(a)
    return strings.Join(a, "")    
}