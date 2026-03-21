package wordy
import (
	"strings"
    "strconv"
)
func Answer(question string) (int, bool) {
    if !strings.HasPrefix(question, "What"){
        return 0, false
    }
    question = strings.ReplaceAll(question," by", "")
	q:= strings.Fields(strings.TrimPrefix(strings.TrimSuffix(question, "?"), "What is "))
    if len(q)%2 == 0 {
        return 0, false
    }
    res, err := strconv.Atoi(q[0])
    if err!= nil{
        return 0, false
    }
    for i:=1;i<len(q); i+=2{
        op:= q[i]
        num, err:= strconv.Atoi(q[i+1])
        if err != nil{
            return 0, false
        }
        switch op{
            case "plus":
            	res+= num
            case "minus":
            	res-= num
            case "multiplied":
            	res*= num
            case "divided":
            	res/=num
        }        
    }
    return res, true   
}