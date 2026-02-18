package luhn

func Valid(id string) bool {
	runes:= []rune{}
    for _, s:= range id{
        if s == ' '{
            continue
        }
        if s <'0' || s> '9'{
            return false
        }
        runes = append(runes, s)
    }
    if len(runes)== 1 && runes[0] == '0'{
        return false
    }
    
    digit:= make([]int, len(runes))
    for i,v:= range runes{
        digit[i] = int(v-'0')
    }
    
    for i:= len(runes)-2; i>=0;i-=2{
        digit[i]*=2
        if digit[i]>9{
            digit[i]-=9
        }
    }
    sum:=0
    for _,d:= range digit{
        sum += d
    }
    if sum%10==0{
        return true
    }
    return false
}
