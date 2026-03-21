package tournament

import (
    "io"
    "bufio"
    "fmt"
    "strings"
    "sort"
    "errors"
)
type Team struct{
    name string
    MP int
    W int
    D int
    L int
    P int
}
func Tally(reader io.Reader, writer io.Writer) error {
	teams := make(map[string]*Team)
    getTeam := func(name string) *Team{
        if v,ok:= teams[name]; ok{
            return v
        }
        newTeam := &Team{name: name}
        teams[name] = newTeam
        return newTeam
    }
    scanner := bufio.NewScanner(reader)
    for scanner.Scan(){
        line:= strings.TrimSpace(scanner.Text())
        if line == ""{
            continue
        }
        if strings.HasPrefix(line, "#"){
            continue
        }
        parts:= strings.Split(line, ";")
        if len(parts) != 3{
            return errors.New("not enough data")
        }
        t1, t2, res := getTeam(parts[0]), getTeam(parts[1]), parts[2]
        switch res{
            case "win":
			t1.W++
            t1.P+=3
            t2.L++
            case "loss":
			t2.W++
            t2.P+=3
            t1.L++			
            case "draw":
            t1.D++
            t1.P++
            t2.D++
            t2.P++
            default:
            	return errors.New("invalid result")
        }
        t1.MP++
        t2.MP++       
    }
    table := []*Team{}
    for _,v:= range teams{
        table = append(table, v)
    }
    sort.Slice(table, func(i,j int) bool{
        if table[i].P != table[j].P{
            return table[i].P > table[j].P
        }
        return table[i].name < table[j].name
    })
    _, err:= fmt.Fprintln(writer,"Team                           | MP |  W |  D |  L |  P")
    if err != nil{
        return err
    }
    for _, v:= range table{
            if _, err:= fmt.Fprintf(writer,"%-31s| %2d | %2d | %2d | %2d | %2d\n", v.name, v.MP, v.W, v.D, v.L, v.P);err != nil{
        return err
    }
    }
    return nil    
}