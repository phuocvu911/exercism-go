package prime
import "errors"
// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n<1 {
        return 0, errors.New("there is no prime for the input")
    }
    count:= 0
    num:= 1
    for count< n{
        num++
        if isPrime(num){
            count++
        }
    }
    return num, nil
}
func isPrime(n int) bool{
    if n<2 {
        return false
    }
    for i:=2; i*i<=n; i++{ //we only need to check until squareroot of n, stop early.
        if n%i==0{
            return false
        }
    }
    return true
}