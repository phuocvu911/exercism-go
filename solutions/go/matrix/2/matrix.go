package saddlepoints
import (
    "strings"
    "errors"
    "strconv"
)
// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	check := strings.Split(s, "\n")
    rows:= make([][]int, len(check))
    rowlen := 0
    for i, r:= range check{
        num := strings.Fields(r)
        if i==0{
            rowlen = len(num)
        } else if rowlen != len(num){
            return nil, errors.New("rows are not equal")
        }
        rows[i] = make([]int, rowlen)
        for j, n:= range num{
            a, err := strconv.Atoi(n)
            if err != nil{
                return nil, err
            }
            rows[i][j] =a   
        }
    }
    return Matrix(rows), nil    
}

// Cols and Rows must return the results without affecting the matrix. Because slice is ptr to array
func (m Matrix) Cols() [][]int {
	cols:= make([][]int, len(m[0]))
    for i:= range cols{
        cols[i] = make([]int, len(m))
        for j:= range len(m){
            cols[i][j] = m[j][i]
        }
    }
    return cols
}

func (m Matrix) Rows() [][]int {
	rows:= make([][]int, len(m))
    for i:= range rows{
        rows[i] = make([]int, len(m[0]))
        copy(rows[i], m[i])
    }
    return rows	
}

func (m Matrix) Set(row, col, val int) bool {
	if row >= len(m) || col >= len(m[0]) || row < 0 || col < 0{
        return false
    }
    m[row][col] = val
    return true
}
