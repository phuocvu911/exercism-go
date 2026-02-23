package prime

func Factors(n int64) []int64 {
	var res []int64
    var divisor = int64(2)
    for n>1{
        for n%divisor ==0{
            res = append(res, divisor)
            n/= divisor
        }
        divisor ++
    }
    return res
}
