package allyourbase
import (
    "math"
    "errors"
    )

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase <2{
        return nil, errors.New("input base must be >= 2")
    }
    if outputBase <2{
        return nil, errors.New("output base must be >= 2")
    }
	num := 0
    for i, v:= range inputDigits{
        if v < 0 || v >= inputBase{
            return nil, errors.New("all digits must satisfy 0 <= d < input base")
        }
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