package transpose
import "strings"
func Transpose(input []string) []string {
	maxLen:= 0
    for _,v := range input{
        if maxLen<len(v){
            maxLen = len(v)
        }
    }
    res:= make([]string, maxLen)
   	for i:=range maxLen{
        var sb strings.Builder
        for j:= range input{
            if i < len(input[j]){
                sb.WriteByte(input[j][i])
            } else {
                hasLetter := false
                for k:= j+1; k< len(input); k++{
                    if i < len(input[k]){
                        hasLetter = true
                        break
                    }
                }
                if hasLetter{
                    sb.WriteByte(' ')
                }               
            }
        }
        res[i] = sb.String()
    }
    return res
}