package atbash
import "strings"
import "regexp"

func Atbash(s string) string {
	s = strings.ToLower(s)
    re:= regexp.MustCompile(`[a-z0-9]+`)
    words:= re.FindAllString(s, -1)
    x := strings.Join(words, "")
    res := ""
    for _, c := range x{
        if c>= 'a' && c<= 'z'{
            res+= string('z'-(c-'a'))
        } else {
            res += string(c)
        }
    }
    a:= []rune(res)
    b:= ""
    if len(a) >= 5{
        for len(a)>5{
        	b+= string(a[:5])
            b+= " "
            a = a[5:]
        }
        b+= string(a)
        return b
    } else {
        return res
    }
}
