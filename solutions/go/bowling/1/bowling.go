package bowling
import "errors"
// Define the Game type here.
type Game struct{
    rolls [21]int
    currentRoll int
    currentFrame int
    done bool
}
func NewGame() *Game {
	return &Game{}
}

func (g *Game) Roll(pins int) error {
	if g.done {
        return errors.New("game is over")
    }
    if pins < 0 || pins > 10{
        return errors.New("invalid roll")
    }
    if g.currentFrame < 9 {
        first:= g.rolls[g.currentFrame*2]
        if first + pins > 10{
            return errors.New("Pin count exceeds pins on the lane")
        }    
    	g.rolls[g.currentFrame*2 + g.currentRoll] = pins    
        if pins == 10 || g.currentRoll == 1{
            g.currentFrame++
            g.currentRoll = 0
        } else {
            g.currentRoll = 1
        }
    } else if g.currentFrame == 9 {
        first:= g.rolls[18]
        second:= g.rolls[19]
        if (first + pins > 10 && first < 10 && g.currentRoll == 1){
            return errors.New("Pin count exceeds pins on the lane")
        }
        if (first ==10 && second + pins > 10 && second < 10 && g.currentRoll == 2){
            return errors.New("Pin count exceeds pins on the lane")
        }
        
        g.rolls[g.currentFrame*2 + g.currentRoll] = pins
        second= g.rolls[19]
        if g.currentRoll == 1 && first != 10 && first + second != 10{
            g.done = true //open frame in Frame 10
        } else if g.currentRoll == 2{
            g.done = true
        }
        g.currentRoll++
    }
    return nil
}

func (g *Game) Score() (int, error) {
	if !g.done{
        return 0, errors.New("incomplete game")
    }
    score := 0
    for f:= 0; f<9; f++{
        i:=2*f
        if g.rolls[i] == 10{
            if g.rolls[2*(f+1)] == 10 && f != 8{
                score += 10 + 10 + g.rolls[2*(f+2)]
            } else {
                score += 10 + g.rolls[i+2] + g.rolls[i+3]
            }         
        } else if g.rolls[i] + g.rolls[i+1] == 10{
            score += 10 + g.rolls[i+2]
        } else {
            score += g.rolls[i] + g.rolls[i+1]
        }
    }
    score += g.rolls[18] + g.rolls[19] + g.rolls[20]
    return score, nil
}