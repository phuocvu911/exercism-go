package stateoftictactoe
import "errors"
var Err = errors.New("invalid")
type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(b []string) (State, error) {
    var countX, countO int
    for i:= range b{
        for _, v:= range b[i]{
            if v == 'X'{
                countX++
            }
            if v == 'O'{
                countO++
            }
        }
    }
    if countO > countX || countX-countO > 1{
        return "", Err
    }
    xWin:= isWinner(b, 'X')
    oWin:= isWinner(b, 'O')
    if xWin && oWin{
        return "", Err
    }
    if xWin && countX != countO +1{
        return "", Err
    }
    if oWin && countX != countO{
        return "", Err
    }
    if xWin || oWin {
        return Win, nil
    }
    if countX + countO == 9{
        return Draw, nil
    }
    return Ongoing, nil
}

func isWinner(b []string, p byte) bool{
    line := func(x,y,z byte) bool{
        return x==p&&y==p&&z==p
    }
    for i:= range b{
        if line(b[i][0], b[i][1], b[i][2]) || line(b[0][i], b[1][i], b[2][i]){
            return true
        }
    }
    return line(b[0][0], b[1][1], b[2][2]) || line(b[0][2], b[1][1], b[2][0])
}
