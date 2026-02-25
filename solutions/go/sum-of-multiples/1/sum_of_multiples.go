package summultiples
import "sort"
func SumMultiples(limit int, divisors ...int) int {
	if len(divisors) == 0{
        return 0
    }
    res:= []int{}
    
    for _, v:= range divisors{
        a:= multi(limit, v)
        res= append(res, a...)
    }
    sort.Slice(res, func(i, j int) bool{
        return res[i] < res[j]
    })
    if len(res) == 0{
        return 0
    }
    sum := res[0]
    for i:= 1; i< len(res); i++{
        if res[i] != res[i-1]{
            sum+=res[i]
        }
    }
    return sum
}



func multi(limit int, div int) []int{
    res:=[]int{}
    if div == 0{
        return res
    }
    for i:=div; i<limit; i+=div{
        res = append(res,i)
    }
    return res
}