package cryptosquare
import(
    "strings"
    "regexp"
    "math"
)
func Encode(pt string) string {
	re:= regexp.MustCompile(`[a-z0-9]+`)
    s:= strings.Join(re.FindAllString(strings.ToLower(pt),-1),"")
    l:= len(s)
    c:= int(math.Ceil(math.Sqrt(float64(l))))//c=8 0 8 16 1 9 17
    r:= int(math.Floor(math.Sqrt(float64(l))))//r=7
    if c*r < l{
        r++
    }
	var res strings.Builder
    for col:= 0; col<c; col++{
        for row:=0; row<r; row++{
            index:= row*c + col
            if index < l{
                res.WriteByte(s[index])
            } else{
                res.WriteByte(' ')
            }
        }
        if col < c-1{
            res.WriteByte(' ')
        }
    }
    return res.String()   
}