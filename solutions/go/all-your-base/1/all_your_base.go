package allyourbase
import (
    "math"
    "errors"
    )
var Err = errors.New("input base must be >= 2")
var Err1 = errors.New("output base must be >= 2")
var Err2 = errors.New("all digits must satisfy 0 <= d < input base")
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase <2{
        return nil, Err
    }
    if outputBase <2{
        return nil, Err1
    }
    for _, v:= range inputDigits{
        if v <0 || v>= inputBase{
            return nil, Err2
        }
    }
	num:= 0
    for i, v:= range inputDigits{
        num += v*int(math.Pow(float64(inputBase), float64(len(inputDigits)-1-i)))
    }
    if num==0{
        return []int{0}, nil
    }
    res:= []int{}
    for num>0{
        digit := num%outputBase
        res = append([]int{digit}, res...) //prepend technique
        num/=outputBase
    }
    return res, nil
}