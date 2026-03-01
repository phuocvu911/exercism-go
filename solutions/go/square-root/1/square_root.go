package squareroot
import "errors"
func SquareRoot(number int) (int, error) {
	var pow func(n int) bool
    pow = func(n int) bool{
        if n*n == number{
            return true
        }
        return false
    }
    for i:=1; i<=number; i++{
        if pow(i){
            return i, nil
        }
    }
    return 0, errors.New("err")
}
