package eliudseggs

func EggCount(n int) int {
	res:=0
    for n>0{
        if n%2 == 1{
            res++
        }
        n/=2
    }
    return res
}
