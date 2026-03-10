package saddlepoints
import (
    "strings"
    "errors"
    "strconv"
)
// Define the Matrix type here.
type Matrix struct{
    rows [][]int
    cols [][]int
}
func New(s string) (Matrix, error) {
	check := strings.Split(s, "\n")
    rows:= make([][]int, len(check))
    rowlen := 0
    for i, r:= range check{
        num := strings.Fields(r)
        if i==0{
            rowlen = len(num)
        } else if rowlen != len(num){
            return Matrix{nil,nil}, errors.New("rows are not equal")
        }
        rows[i] = make([]int, rowlen)
        for j, n:= range num{
            a, err := strconv.Atoi(n)
            if err != nil{
                return Matrix{nil,nil}, err
            }
            rows[i][j] =a   
        }
    }
    cols := make([][]int, len(rows[0]))
    for i:= range cols{
        cols[i] = make([]int, len(rows))
        for j:= range rows{
            cols[i][j] = rows[j][i]
        }
    }
    return Matrix{rows, cols}, nil    
}

// Cols and Rows must return the results without affecting the matrix. Because slice is ptr to array
func (m Matrix) Cols() [][]int {
	cols:= make([][]int, len(m.cols))
    for i:= range cols{
        cols[i] = make([]int, len(m.cols[0]))
        copy(cols[i], m.cols[i])
    }
    return cols
}

func (m Matrix) Rows() [][]int {
	rows:= make([][]int, len(m.rows))
    for i:= range rows{
        rows[i] = make([]int, len(m.rows[0]))
        copy(rows[i], m.rows[i])
    }
    return rows	
}

func (m *Matrix) Set(row, col, val int) bool {
	if row >= len(m.rows) || col >= len(m.cols) || row < 0 || col < 0{
        return false
    }
    m.rows[row][col] = val
    m.cols[col][row] = val
    return true
}
