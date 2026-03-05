package gameoflife

func Tick(matrix [][]int) [][]int {
    if len(matrix) == 0{
        return matrix
    }
	row, col:= len(matrix), len(matrix[0])
    dirs:= [][]int{
        {-1,-1}, {-1,0}, {-1,1},
        {0,-1}	/*X*/  , {0,1},
        {1,-1}, {1,0}  ,{1,1},
    }
    res:= make([][]int, row)
    for i,_:= range res{
        res[i] = make([]int, col)
    }
    for i:=0; i<row; i++{
        for j:=0;j<col; j++{
            count := 0
            for _, d:= range dirs{
                x:= i+d[0]
                y:= j+d[1]
                if x>=0 && x< row && y>= 0 && y < col{
                    if matrix[x][y] == 1{
                        count++ //count adjacent live
                    }
                }
            }
            if matrix[i][j] == 0 && count == 3{
                res[i][j] = 1
            } else if matrix[i][j] == 1{
                if count == 2 || count == 3{
                    res[i][j] = 1
                } 
            } 
        }
    }
    return res
}
