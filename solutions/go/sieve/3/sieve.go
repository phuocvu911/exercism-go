package sieve
func Sieve(limit int) []int {
    res:= []int{}
	p := make([]bool, limit-1)  
    for i:= 0; i<limit -1; i++{
        if p[i]{
            continue
        } else{
            res = append(res, i+2)
            for j:= i+i+2; j<limit-1; j+=i+2{
                p[j] = true
            }
        }
    }
    return res  
}