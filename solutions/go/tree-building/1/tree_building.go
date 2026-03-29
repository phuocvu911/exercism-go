package tree
import "errors"
import "sort"
type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	n:= len(records)
    if n == 0{
        return nil, nil
    }
    nodes:= make([]*Node, n)
    for i:= range n{
        nodes[i] = &Node{ID:i}
    }
    var root *Node
    sort.Slice(records, func(i,j int)bool{
        return records[i].ID<records[j].ID
    })
    rootCount:=0
    for _, r:= range records{
		if r.ID < 0 || r.ID >= n {
			return nil, errors.New("invalid ID")
		}
		if r.Parent < 0 || r.Parent >= n {
			return nil, errors.New("invalid ID")
		}
		if r.Parent >= r.ID && r.Parent !=0{
			return nil, errors.New("invalid ID")
        }        
        node:= nodes[r.ID]       
        if r.ID == 0 && r.ID == r.Parent{
            root = node
            rootCount++
        } else {
        	parent:= nodes[r.Parent]
            parent.Children = append(parent.Children, node)
        }
    }
    seen := make([]bool, n)
    for _, r := range records {
        if seen[r.ID] {
            return nil, errors.New("duplicate ID")
        }
        seen[r.ID] = true
    }
    if rootCount != 1{
        return nil, errors.New("no root found or too many roots")
    }
    return root, nil
}