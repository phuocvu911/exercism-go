package camicia
import "strings"
type Outcome struct {
	finishes bool
	cards    int
	tricks   int
}
var penaltyMap = map[string]int{
    "J":1,
    "Q":2,
    "K":3,
    "A":4,
}

func SimulateGame(playerA, playerB []string) Outcome {
	A:= append([]string{}, playerA...)
	B:= append([]string{}, playerB...)
    var pile []string
    current, penalty:=0,0 //0 for A, 1 for B turn
    lastFace:= -1
    seen:= make(map[string]bool)
    cards, tricks:= 0,0
    for{
        //check state for loop detection(only when there is no penalty)
        if penalty == 0{
            state:= getState(A) + "|" + getState(B) + "|" +string(rune(current))
        if seen[state]{
            return Outcome{false, cards, tricks}
        }
        seen[state] = true}

        //check finnish (when its a player turn and they have no card)
        if current == 0 && len(A) == 0{
            if len(pile) > 0{
                B = append(B, pile...)
                tricks++
            }
            return Outcome{true, cards, tricks}
        }
        if current == 1 && len(B) == 0{
            if len(pile) > 0{
                A = append(A, pile...)
                tricks++
            }            
            return Outcome{true, cards, tricks}
        }

        //draw card
        var card string
        if current == 0 {
            card = A[0]
            A = A[1:]
        } else{
            card = B[0]
            B = B[1:]
        }
		pile = append(pile, card)
        cards++

        //check facecard
        if p, ok:= penaltyMap[card]; ok{
            penalty = p
            lastFace = current
            current ^= 1
            continue
        }

        //paying penalty
        if penalty >0 {
            if (current == 0 && len(A) == 0) || (current == 1 && len(B) == 0){
            	if lastFace == 1{
                	B = append(B, pile...)
            	} else{
                    A= append(A, pile...)
                }
                tricks++
            	return Outcome{true, cards, tricks}
        	}
            penalty--
            if penalty ==0{
                if lastFace == 0 {
                    A = append(A, pile...)
                } else{
                    B = append(B, pile...)
                }
                tricks++
                pile = nil
                current = lastFace
            }
        } else{
            current ^= 1
        }      
    }
}

func getState(s []string) string{
    var res strings.Builder
    for _, v:= range s{
        if _, ok:= penaltyMap[v]; ok{
            res.WriteString(v)
        } else{
            res.WriteString("N")
        }
    }
    return res.String()
}
