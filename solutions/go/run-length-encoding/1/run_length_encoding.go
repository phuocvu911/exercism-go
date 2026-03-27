package runlengthencoding
import (
    "strconv"
    "regexp"
)

func RunLengthEncode(input string) string {
	var letter []byte
	var count []int   
    if input == ""{
        return ""
    }
    letter = append(letter, input[0])
    cnt:= 1
    for i:=1; i<len(input); i++{      
        if input[i] == input[i-1]{
            cnt++
        } else {
            count = append(count, cnt)
            cnt = 1
            letter = append(letter, input[i])
        }
    }
    count = append(count, cnt)
    res :=""
    for i,v:= range letter{
        if count[i] != 1{
            res+= strconv.Itoa(count[i])
        }
        res+= string(v)
    }
    return res
}

func RunLengthDecode(input string) string {
    if input == ""{
        return ""
    }
    re:= regexp.MustCompile(`[a-zA-Z\s]+`)
    letter:= re.FindAllString(input,-1)
    count:= re.Split(input, -1)
    count= count[:len(count)-1]
    res := ""
    for i,v:= range letter{
        num, _ := strconv.Atoi(count[i])
        for j:=0; j< num-1; j++{
            res += string(v[0])
        }
        res += v
    }
    return res
}
