package eliudseggs

func EggCount(n int) int {
	res:=0
    for n>0{
        digit:= n%2
        if digit == 1{
            res++
        }
        n/=2
    }
    return res
}
