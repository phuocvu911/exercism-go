package lsproduct
import(
    "errors"
    "strconv"
)
func LargestSeriesProduct(s string, span int) (int64, error) {
    if span < 0{
        return 0, errors.New("span must not be negative")
    }
    if span == len(s){
		return Product(s)
    }    
    if span > len(s){
        return 0, errors.New("span must be smaller than string length")
    }
    res, _:= Product(s[0:span])
    for i:= 1; i<len(s)-span+1; i++{
        b, err:= Product(s[i:i+span])
        if err != nil{
            return 0, err
        }
        if res < b{
            res = b
        }
    }
    return res, nil
}
func Product(s string) (int64, error) {
    res := int64(1)
	for _, v:= range s{
        num, err:= strconv.Atoi(string(v))
        if err != nil{
            return 0, errors.New("invalid number")
        }
        res *= int64(num)
    }
    return res, nil
}