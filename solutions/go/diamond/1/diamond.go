package diamond
import(
    "strings"
    "errors"
)
func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z'{
        return "", errors.New("input has to be uppercase letter from A-Z")
    }
    pos:= int(char-'A')
    size := pos*2+1
    var res strings.Builder
    for i:=0; i<size; i++{
        var idx int
        if i<=pos{
            idx=i
        } else{
            idx= size - i -1
        } //mirroring
        
		letter:= byte(idx+'A')
        outter:= strings.Repeat(" ", pos-idx)
        res.WriteString(outter)
        if idx ==0{
            res.WriteByte(letter)
        } else{
            inner:= strings.Repeat(" ", 2*idx-1) //first and last row we dont need inner spaces
            res.WriteByte(letter)
			res.WriteString(inner)
            res.WriteByte(letter)
        }
        res.WriteString(outter)
        if i<size-1{
            res.WriteByte('\n')
        }
    }
    return res.String(), nil
}