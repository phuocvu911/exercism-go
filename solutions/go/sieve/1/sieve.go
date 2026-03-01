package sieve

func Sieve(limit int) []int {
    res:= []int{}
    if limit < 2{
        return res
    }
	h:= make([]int, limit-1)
    for i:= 0;i< limit-1; i++{
        h[i] = i+2
    }
    
    p:= make([]bool, limit-1)
    for i, v:= range h{
        if p[i]{
            continue
        } else if !p[i]{
            res = append(res, v)
            for j:= i+v; j<limit-1; j+=v{
                p[j] = true
            }
        }
    }
    return res  
}