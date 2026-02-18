package isbn
import "strings"
func IsValidISBN(s string) bool {
	s = strings.ReplaceAll(s, "-","")
    sum:= 0
    x:=10
	for _, v := range s{
        if x == 1 && s[9] == 'X'{
            sum += 10
            break
        }
        if v<'0' || v>'9'{
            return false
        }
        sum += int(v-'0') * x
        x--
    }
	return sum%11==0 && len(s) == 10
}