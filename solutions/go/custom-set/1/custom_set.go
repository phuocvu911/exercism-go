package customset
import "fmt"

type Set struct{
    data map[string]struct{} //key in map is unique, empty struct is using no memory
}
func New() Set {
	return Set{data:make(map[string]struct{})}
}

func NewFromSlice(l []string) Set {
	n:= New()
    for _, v:= range l{
        n.data[v] = struct{}{} //empty value of struct
    }
    return n
}

func (s Set) String() string {
    if len(s.data)==0{
        return "{}"
    }
    res:="{"
    for key,_:= range s.data{
        res+= fmt.Sprintf("\"%s\"", key)
        res += ", "
    }
    res= res[:len(res)-2]
    res+="}"
    return res
}

func (s Set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s Set) Has(elem string) bool {
	_, ok:= s.data[elem]
    return ok
}

func (s Set) Add(elem string) {
	s.data[elem] = struct{}{}
}
// test if s1 is subset of s2, we check if s2 has all element of s1
func Subset(s1, s2 Set) bool {
	for k:= range s1.data{
        if !s2.Has(k){
            return false
        }
    }
    return true
}

func Disjoint(s1, s2 Set) bool {
	for k:= range s1.data{
        if s2.Has(k){
            return false
        }
    }
    return true	
}

func Equal(s1, s2 Set) bool {
	return len(s1.data) == len(s2.data) && Subset(s1,s2)
}

func Intersection(s1, s2 Set) Set {
	res:= New()
    for k:= range s1.data{
        if s2.Has(k){
            res.Add(k)
        }
    }
    return res
}

func Difference(s1, s2 Set) Set {
	res:= New()
    for k:= range s1.data{
        if !s2.Has(k){
            res.Add(k)
        }
    }
    return res
}

func Union(s1, s2 Set) Set {
	for k:= range s1.data{
        s2.Add(k)
    }
    return s2
}