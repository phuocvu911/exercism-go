package wordy
import (
	"strings"
    "strconv"
)
type opFunc func(int, int) int
var opMap = map[string]opFunc{
    "plus": func(a,b int) int{return a+b},
    "minus": func(a,b int) int{return a-b},
    "mul": func(a,b int) int{return a*b},
    "div": func(a,b int) int{return a/b},
}
func Answer(question string) (int, bool) {
    if !strings.HasPrefix(question, "What"){
        return 0, false
    }
    question = strings.ReplaceAll(question,"multiplied by", "mul")
    question = strings.ReplaceAll(question,"divided by", "div")
	q:= strings.Fields(strings.TrimPrefix(strings.TrimSuffix(question, "?"), "What is "))
    if len(q) == 0 || len(q) == 2{
        return 0, false
    }
    res, err := strconv.Atoi(q[0])
    if err!= nil{
        return 0, false
    }
    for i:=1;i<len(q); i+=2{
        s := q[i]
        op, ok:= opMap[s]
        if !ok{
            return 0, false
        }
        numStr := q[i+1]
        num, err:= strconv.Atoi(numStr)
        if err != nil{
            return 0, false
        }
        res = op(res, num)
    }
    return res, true   
}