package queenattack
import (
    "regexp"
    "errors"
    "strconv"
)
func CanQueenAttack(p1,p2 string) (bool, error) {
	re:= regexp.MustCompile(`^[a-h][1-8]$`)
    if !re.MatchString(p1) || !re.MatchString(p2) {
        return false, errors.New("invalid")
    }
    if p1 == p2{
        return false, errors.New("same ")
    }
    col1, col2 := p1[0], p2[0]
    row1, _ := strconv.Atoi(string(p1[1]))
    row2, _ := strconv.Atoi(string(p2[1]))
    if col1 == col2 || row1 == row2 || int(col1)-int(col2) == row1-row2 || int(col2) - int(col1) == row1-row2{
        return true, nil
    }
    return false, nil
}
