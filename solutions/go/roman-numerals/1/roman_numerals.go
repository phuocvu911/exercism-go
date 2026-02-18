package romannumerals
import "errors"
func ToRomanNumeral(n int) (string, error) {
	if n<1 || n>3999{
        return "", errors.New("err")
    }
// 2 slices is prefered than a map, need order 
	values:= []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
    romans:= []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
    
    res:= ""
    for i,v := range values{
        for n>=v{
            res+= romans[i]
            n -= v
        }
    }
    return res, nil   
}