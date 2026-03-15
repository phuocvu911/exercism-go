package perfectnumbers
import "errors"
type Classification int
var ErrOnlyPositive = errors.New("input must be positive")
const (
	ClassificationDeficient Classification = 1
	ClassificationPerfect Classification =  2
	ClassificationAbundant Classification = 3
)

func Classify(n int64) (Classification, error) {
	if n<= 0 {
        return 0, ErrOnlyPositive
    }
    sum := Sum(n)
    switch{
        case sum == n:
            return ClassificationPerfect, nil
        case sum < n:
        	return ClassificationDeficient, nil
        default:
        	return ClassificationAbundant, nil
    }   
}
func Sum(n int64) int64{
    sum:=int64(0)
    if n==1{
        return sum
    }
    for i:= int64(1); i*i<= n; i++{
        if n%i==0{
            sum += i
            if i != n/i && i!= 1{
                sum += n/i
            }
        }
    }
    return sum  
}
