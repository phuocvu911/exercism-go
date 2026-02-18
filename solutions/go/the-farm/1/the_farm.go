package thefarm
import "errors"
import "fmt"
/*
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}*/
// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, cows int) (float64, error){
	res, err := f.FodderAmount(cows)
    if err != nil{
        return 0.0, err
    }
    res1, err1 := f.FatteningFactor()
    if err1 != nil{
        return 0.0, err1
    }
    return res*res1/float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error){
    if cows <= 0{
        return 0.0, errors.New("invalid number of cows")
    }

    return DivideFood(f, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct{
    Cows int
    Msg string
}

func (i *InvalidCowsError) Error() string{
    return fmt.Sprintf("%d cows are invalid: %s", i.Cows, i.Msg)
}

func ValidateNumberOfCows(cows int) error{
    if cows < 0{
        return &InvalidCowsError{
            cows,
            "there are no negative cows",
        }
    } else if cows == 0{
        return &InvalidCowsError{
            cows,
            "no cows don't need food",
        }
    }

    return nil
}

