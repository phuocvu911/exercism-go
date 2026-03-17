package flowerfield
var dirs = [][]int{//dirs is 8 directions around each position in board
        {-1,-1},{-1,0},{-1,1},
        {0,-1},	/*X*/  {0,1},
        {1,-1}, {1,0}, {1,1},
    }
// Annotate returns an annotated board
func Annotate(board []string) []string {
	if len(board) == 0{
        return board
    }
    col:= len(board[0])
    row:= len(board)
    field:= make([][]rune, row)
    for i, v:= range board{
        field[i] = []rune(v)
    }
    
    for i:= range row{
        for j:= range col{
            if field[i][j] != ' '{
                continue //skip the * position
            }
            count:=0
            for _, d:= range dirs{
                x:= i+d[0]
                y:=j+d[1] //x,y is position of the adjacent pos
                if x >=0 && x<row && y>=0 && y<col{
                    if field[x][y] == '*'{
                        count++
                    }
                }
            }
            if count>0{
                field[i][j] = rune(count + '0')
            }
        }
    }
    res:= make([]string, row)
    for i,v:= range field{
        res[i] = string(v)
    }
    return res
}