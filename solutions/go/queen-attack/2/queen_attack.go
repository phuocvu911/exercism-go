package queenattack
import (
    "regexp"
    "errors"  
)
func CanQueenAttack(p1,p2 string) (bool, error) {
	re:= regexp.MustCompile(`^[a-h][1-8]$`)
    if !re.MatchString(p1) || !re.MatchString(p2) || p1 == p2{
        return false, errors.New("invalid")
    }
    col1, col2, row1, row2 := p1[0], p2[0], p1[1], p2[1]
    return col1 == col2 || row1 == row2 || col1-col2 == row1-row2 || col2 - col1 == row1-row2, nil
}