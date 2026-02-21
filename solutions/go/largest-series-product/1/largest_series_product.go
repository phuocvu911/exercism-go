package lsproduct
import(
    "errors"
    "strconv"
)
func LargestSeriesProduct(s string, span int) (int64, error) {
    var res int64
    res = 1
    if span < 0{
        return 0, errors.New("span must not be negative")
    }
    if span == len(s){
		return Product(s)
    }
    
    if span > len(s){
        return 0, errors.New("span must be smaller than string length")
    }
    s+="0"
    serie1 := s[0:span]
    res, _= Product(serie1)
    for i:= 1; i<len(s)-span; i++{
        serie2 := s[i:i+span]
        b, err:= Product(serie2)
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
    var res int64
    res = 1
	for _, v:= range s{
        num, err:= strconv.Atoi(string(v))
        if err != nil{
            return 0, errors.New("invalid number")
        }
        res *= int64(num)
    }
    return res, nil
}
