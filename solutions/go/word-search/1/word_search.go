package wordsearch
import "errors"
var dirs=[][]int{//dirs is 8 directions around each position in board
        {-1,-1},{-1,0},{-1,1},
        {0,-1},	/*X*/  {0,1},
        {1,-1}, {1,0}, {1,1},
    }
func Solve(words []string, puzzle []string) (res map[string][2][2]int, err error) {
	if len(puzzle) == 0{
        return nil, errors.New("empty search area")
    }
	res = make(map[string][2][2]int, len(words))
    row:= len(puzzle)
    col:= len(puzzle[0])
    for _, w:= range words{
        found:= false
        for r:= range row{
            for c:= range col {
                for _,d:= range dirs{
                    if match(puzzle, w,r,c,d[0], d[1]){
                        endR:=r + (len(w)-1)*d[0]
                        endC:=c + (len(w)-1)*d[1]
                        res[w] = [2][2]int{{c,r}, {endC, endR}}
                        found = true
                        break
                    }
                }
            }
        }
        if !found{
            res[w] =[2][2]int{{-1,-1}, {-1,-1}}
            err = errors.New("cant find this word")
        }
    }
    return res, err
}

func match(puzzle []string, word string, r,c,dr,dc int) bool{
    row:= len(puzzle)
    col:= len(puzzle[0])
    for i:= range word{
        nr:= r+ i*dr
        nc:= c+ i*dc
        if nr<0 || nr >= row || nc < 0 || nc >= col{ // out bound
            return false
        }
        if word[i] != puzzle[nr][nc]{ //word not match
            return false
        }
    }
    return true
}