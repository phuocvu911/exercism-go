package spiralmatrix
//left, down, right, up
func SpiralMatrix(size int) [][]int {
	res:= make([][]int, size)
    for i:=range res{
        res[i] = make([]int, size)
    }
    top, bottom:= 0, size-1
	left, right:= 0, size-1
    n:=1
    for n<= size*size{
        for i:= left; i<= right; i++{
            res[top][i] = n
            n++
        }
        top++
        for i:=top; i<= bottom; i++{
            res[i][right] = n
            n++
        }
        right--
        for i:=right; i>= left; i--{
            res[bottom][i] = n
            n++
        }
        bottom--
        for i:= bottom; i>= top; i--{
            res[i][left] = n
            n++
        }
        left++
    }
    return res
}