package hangman
import(
    "errors"
    "strings"
)
type Game struct{
    word string
    guess map[rune]bool
    guessleft int
    state string
}

func NewGame(word string) *Game {
    return &Game{
        word: word,
        guess: make(map[rune]bool),
        guessleft: 9,
        state: "Ongoing",
    }
}

func (g *Game) Guess(r rune) error {
	if g.state == "Lose"{
        return errors.New("cannot guess after the game is lost")
    }
    if g.state == "Win"{
        return errors.New("cannot guess after the game is won")
    }
    if g.guess[r] || !strings.ContainsRune(g.word, r){
        if g.guessleft >0{
            g.guessleft--
        } else{
            g.state = "Lose"
        }
    }    
    g.guess[r] = true
    if g.MaskedWord() == g.word{
        g.state = "Win"
    }
    return nil    
}

func (g *Game) MaskedWord() string {
	var res strings.Builder
    for _, r:= range g.word{
        if g.guess[r] {
            res.WriteRune(r)
        }else{
            res.WriteRune('_')
        }
    }
    return res.String()
}

func (g *Game) RemainingGuesses() int {
	return g.guessleft
}

func (g *Game) State() string {
	return g.state
}