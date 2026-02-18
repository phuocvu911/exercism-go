package grains
import "errors"
func Square(n int) (uint64, error) {
	if n<=0 || n>64{
        return 0, errors.New("err") 
    }
    return 1<<(n-1), nil // bit-shifting
}

func Total() uint64 {
    return 1<<64 -1 // sum of all bit-shifting from 1-64, or shift n=65 -1
}
