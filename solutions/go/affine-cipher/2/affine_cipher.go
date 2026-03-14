package affinecipher
import (
    "errors"
    "strings"
)
const m = 26
func Encode(text string, a, b int) (string, error) {
	if !gcd(a, m){
        return text, errors.New("a and m are not coprime")
    }
    res:= []rune{}
    count := 0
    for _, r:= range strings.ToLower(text){
        if (r<'a' || r>'z') && (r<'0' || r>'9'){ //exclude punct
            continue
        } else{            
            if count>0 && count%5 == 0{
                res = append(res, ' ')
            }
            if r>='0' && r<='9'{ //keep digit
                res = append(res,r)
            } else{
            	res = append(res, rune((a*int(r-'a')+b)%m + 'a')) //transform letter
            }
            count++            
        }
    }
    return string(res), nil
}

func Decode(text string, a, b int) (string, error) {
	if !gcd(a, m){
        return text, errors.New("a and m are not coprime")
    }
    b= b%m
    res:= []rune{}
    x:= mmi(a,m)
    for _, r:= range text{
        if (r<'a' || r>'z') && (r<'0' || r>'9'){ //exclude punct
            continue
        } else {
            if r>='0' && r<='9'{ //keep digit
                res = append(res,r)  
            } else {                
            	y:= int(r-'a')
                pos:= (x*(y-b+m)%m)%m
                res = append(res, rune('a'+pos))
            }            
        }       
    }
    return string(res), nil 
}

func gcd(a,m int) bool{
	for m!=0{
        a,m = m, a%m
    }
    return a == 1
}

func mmi(a,m int) int{
    for x:=1; x<m; x++{
        if a*x%m == 1{
            return x
        }
    }
    return 1
}
