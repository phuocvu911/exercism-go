package wordcount
import "strings"
import "regexp"
type Frequency map[string]int

func WordCount(phrase string) Frequency {
	s:= strings.ToLower(phrase)
    re:= regexp.MustCompile(`[a-z0-9]+('[a-z0-9]+)*`)
    words:= re.FindAllString(s, -1)
    h:= Frequency{}
    for _, v := range words{
        h[v]++
    }
    return h
}
