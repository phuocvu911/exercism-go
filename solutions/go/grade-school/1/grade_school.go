package gradeschool
import "sort"
type School struct{
    g map[int][]string
}

func New() *School {
	return &School{g: make(map[int][]string)} // grade -> students
}

func (s *School) Add(name string, grade int) bool {
   for _, students:= range s.g{ //check the whole school for that name 
       for _, student:= range students{
           if name == student{
               return false
           }
       }
   	}
    s.g[grade] = append(s.g[grade], name)
    sort.Strings(s.g[grade])
    return true
}

func (s *School) Grade(level int) []string {
	return s.g[level]
}

func (s *School) Enrollment() []string {
	res := []string{}
    for i:=1; i<=12; i++{
        res = append(res, s.g[i]...)
    }
    return res
}