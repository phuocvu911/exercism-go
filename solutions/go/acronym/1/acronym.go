package acronym
import(
    "regexp"
    "strings"
)
func Abbreviate(s string) string {
    s= strings.ToUpper(s)
    re:= regexp.MustCompile(`[A-Z]+['A-Z]*`)
    res:= re.FindAllString(s, -1)
    arc:= ""
	for _, str:= range res{
        arc+= string(str[0])
    }
    return arc
}