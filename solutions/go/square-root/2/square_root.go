package squareroot
import "errors"
func SquareRoot(number int) (int, error) {
    for i:=1; i<=number; i++{
        if i*i==number{
            return i, nil
        }
    }
    return 0, errors.New("err")
}
