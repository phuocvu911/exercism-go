package rectangles

func Count(d []string) int {
	rows:= len(d)
    if rows == 0{
        return 0
    }
    cols:= len(d[0])
    count:=0
    for r1:= range rows{
        for c1:= range cols{ //top left corner
            if d[r1][c1] != '+'{
                continue
            }
            for c2:= c1+1; c2< cols; c2++{ // top right corner
                if d[r1][c2] != '+'{
                    continue
                }
                for r2:= r1+1; r2<rows; r2++{
                    if d[r2][c1] != '+' || d[r2][c2] != '+'{ //bottom corners
                        continue
                    }
                    if validHorizon(d,r1,c1,c2) && //top edge
                     validHorizon(d,r2,c1,c2) && // bot edge
                     validVertical(d,c1,r1,r2) && //left edge
                     validVertical(d,c2,r1,r2){ //right edge
                        count++
                    } 
                }
            }
        }
    }
    return count
}

func validHorizon(d []string, row, c1,c2 int) bool{
    for c:= c1+1; c<c2; c++{
        if d[row][c] != '-' && d[row][c] != '+'{
            return false
        }
    }
    return true
}
func validVertical(d []string, col,r1,r2 int) bool{
    for r:= r1+1; r<r2; r++{
        if d[r][col] != '|' && d[r][col] != '+'{
            return false
        }
    }
    return true
}