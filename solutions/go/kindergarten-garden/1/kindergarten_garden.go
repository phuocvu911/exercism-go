package kindergarten
import (
	"strings"
    "errors"
    "slices"
)
type Garden struct{
    row1 string
    row2 string
    position map[string]int
}

var tree = map[byte]string{
    'G': "grass",
    'C': "clover",
    'R': "radishes",
    'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	isValid := strings.HasPrefix(diagram, "\n") && strings.Contains(diagram, "\n")
    if !isValid {
        return nil, errors.New("Invalid diagram")
    }
    rows:= strings.Split(diagram, "\n")
    row1, row2 := rows[1], rows[2]
    if len(row1) != len(row2){
        return nil, errors.New("mismatched rows")
    }
    if len(row1) % 2 != 0 || len(row2) % 2 != 0{
        return nil, errors.New("odd number of cups")
    }
    for _,v := range row1{
        if _, ok := tree[byte(v)]; !ok{
            return nil, errors.New("invalid cup codes")
        }
    }
    for _,v := range row2{
        if _, ok := tree[byte(v)]; !ok{
            return nil, errors.New("invalid cup codes")
        }
    }
    if len(children) > 1{
        for i:=0; i< len(children); i++{
            for j:=i+1; j< len(children); j++{
                if children[i] == children[j]{
                    return nil, errors.New("duplicate name")
                }
            }
        }
    }
    sorted:= make([]string, len(children))
    copy(sorted, children)
    slices.Sort(sorted)
    idx:= make(map[string]int, len(sorted))
    for i,v := range sorted{
        idx[v] = i
    }
    return &Garden{row1, row2, idx}, nil
    
}

func (g *Garden) Plants(child string) ([]string, bool) {  
    idx,ok:= g.position[child]
	if !ok{
        return nil, false
    }
    tree1:= tree[g.row1[2*idx]]
    tree2:= tree[g.row1[2*idx+1]]
    tree3:= tree[g.row2[2*idx]]
    tree4:= tree[g.row2[2*idx+1]]
    res := []string{tree1,tree2,tree3,tree4}
    return res, true
}
