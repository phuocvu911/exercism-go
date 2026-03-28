package saddlepoints
import (
    "strings"
    "errors"
    "strconv"
)

type Matrix [][]int
type Pair [2]int

func (m *Matrix) Saddle() []Pair {
	res:= []Pair{}
    if len(*m) == 0 || len((*m)[0]) == 0{
        return res
    }
    cols := m.Cols()
    for i, row:= range *m{
        max:= row[0]
        for _, r:= range row{
            if max < r{
                max = r
            }
        }
        for j, r:= range row{
            if r==max{
                min:= cols[j][0]
                for _, c:= range cols[j]{
                    if c<min{
                        min =c
                    }
                }
                if min == max{
                    res = append(res, Pair{i+1, j+1})
                }
            }
        }
    }       
    return res
}

func New(s string) (*Matrix, error) {
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
    m := Matrix(rows)
	return &m, nil
}

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