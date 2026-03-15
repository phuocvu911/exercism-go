package pascal

func Triangle(n int) [][]int {
	res:= [][]int{}
    for i:=1; i<=n; i++{
        row:= make([]int, i)
        row[0] = 1
        row[i-1] =1
        if i>2{
            for j:= 1; j<i-1; j++{
                row[j] = res[i-2][j] + res[i-2][j-1]
            }
        }
        res = append(res, row)
    }
    return res
}
