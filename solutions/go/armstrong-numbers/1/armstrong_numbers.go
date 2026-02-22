package armstrong
import "math"
func IsNumber(x int) bool {
    length:= 0
    m:=x
    for m>0{
        length ++
        m/=10
    }
    sum:= 0
    n:= x
    for n>0{
        digit := n%10
        sum+= int(math.Pow(float64(digit), float64(length)))
        n/=10
    }
    return sum == x
}
