package spiralmatrix
var(
    down = [2]int{1,0}
    left = [2]int{0,-1}
    up = [2]int{-1,0}
    right = [2]int{0,1}   
)
func SpiralMatrix(size int) [][]int {
	res:= make([][]int, size)
    if size == 0{
        return res
    }
    for i:=range res{
        res[i] = make([]int, size)
    }
    for i:= 1; i<= size; i++{
        res[0][i-1]=i
    }
    a,b:=0, size-1
    x:= size + 1
    time:= size -1
    for x<=size*size{
        for range time{
            a+=down[0]
            b+=down[1]
            res[a][b] = x
            x++
        }
        for range time{
            a+=left[0]
            b+=left[1]
            res[a][b] = x
            x++            
        }
        time--
        for range time{
            a+=up[0]
            b+=up[1]
            res[a][b] = x
            x++
        }
        for range time{
            a+=right[0]
            b+=right[1]
            res[a][b] = x
            x++
        }
        time--
    }
    return res
}
