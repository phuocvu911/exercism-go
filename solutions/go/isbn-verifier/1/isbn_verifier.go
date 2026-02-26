package isbn
import "strings"
import "strconv"
func IsValidISBN(s string) bool {
	hasHyphen := strings.ContainsRune(s, '-')
    if hasHyphen{
        s = strings.ReplaceAll(s, "-","")
    }
    if len(s) != 10{
        return false
    }
    hasX:= strings.HasSuffix(s, "X")
    res:= 0
    x:=10
    if !hasX{
        for _, v := range s{
            num, err:= strconv.Atoi(string(v))
            if err != nil{
                return false
            }
            res += num*x
            x--
        }
    } else{
        for _, v := range s[:len(s)-1]{
            num, err:= strconv.Atoi(string(v))
            if err != nil{
                return false
            }
            res += num*x
            x--
        }
        res -=1
    }

    if res%11==0{
        return true
    }
    return false
    
}

