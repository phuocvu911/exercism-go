package proverb
import "fmt"
// Proverb should have a comment documenting it.
func Proverb(s []string) []string {
	res:= []string{}
    if len(s) == 0{
        return res
    }
    for i:=0;i<len(s)-1; i++{
        v:= fmt.Sprintf("For want of a %s the %s was lost.", s[i], s[i+1])
        res = append(res, v)
    }
    last:= fmt.Sprintf("And all for the want of a %s.", s[0])
    res = append(res, last)
    return res
}