package grains
import "errors"
import "math"
func Square(n int) (uint64, error) {
	if n<=0 || n>64{
        return 0, errors.New("err") 
    }
    if n ==1{
        return 1, nil
    }
    return uint64(math.Pow(2,float64(n-1))), nil
}

func Total() uint64 {
    var res uint64
	for i:=1; i<=64; i++{
        v,_:= Square(i)
        res+=v
    }
    return res
}
