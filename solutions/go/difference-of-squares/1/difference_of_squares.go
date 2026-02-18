package diffsquares

func SquareOfSum(n int) int {
	sum:=0
    for i:=1; i<=n;i++{
        sum+=i
    }
    return sum*sum
}

func SumOfSquares(n int) int {
	if n==1{
        return 1
    }
    return n*n + SumOfSquares(n-1)
}

func Difference(n int) int {
	x:= SumOfSquares(n) - SquareOfSum(n)
    if x<0{
        return -x
    }
    return x
}
