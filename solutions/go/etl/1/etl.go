package etl
import "strings"
func Transform(in map[int][]string) map[string]int {
	h:= map[string]int{}
    for key,val := range in{
        for _, v:= range val{
            h[strings.ToLower(v)] = key
        }
    }
    return h
}
