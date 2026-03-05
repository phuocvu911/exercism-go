package flowerfield

// Annotate returns an annotated board
func Annotate(board []string) []string {
	if len(board) == 0{
        return board
    }
    col:= len(board[0])
    row:= len(board)
    res:= make([][]rune, row)
    for i, v:= range board{
        res[i] = []rune(v)
    }
    dirs:=[][]int{//dirs is 8 directions around each position in board
        {-1,-1},{-1,0},{-1,1},
        {0,-1},	/*X*/  {0,1},
        {1,-1}, {1,0}, {1,1},
    }
    for i:=0; i< row; i++{
        for j:=0; j<col; j++{
            if res[i][j] != ' '{
                continue //skip the * position
            }
            count:=0
            for _, d:= range dirs{
                x:= i+d[0]
                y:=j+d[1] //x,y is position of the adjacent pos
                if x >=0 && x<row && y>=0 && y<col{
                    if res[x][y] == '*'{
                        count++
                    }
                }
            }
            if count>0{
                res[i][j] = rune(count + '0')
            }
        }
    }
    res1:= make([]string, row)
    for i,v:= range res{
        res1[i] = string(v)
    }
    return res1
}
