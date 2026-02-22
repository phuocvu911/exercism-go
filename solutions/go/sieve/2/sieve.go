package sieve
func Sieve(limit int) []int {
    res:= []int{}
	h,p := make([]int, limit-1), make([]bool, limit-1)
    for i:= 0;i< limit-1; i++{
        h[i] = i+2
    }   
    for i, v:= range h{
        if p[i]{
            continue
        } else{
            res = append(res, v)
            for j:= i+v; j<limit-1; j+=v{
                p[j] = true
            }
        }
    }
    return res  
}