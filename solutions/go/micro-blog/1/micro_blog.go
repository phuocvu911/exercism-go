package microblog

func Truncate(s string) string {
	if len(s)<=5{
        return s
    }
    res:= []rune(s)    
    return string(res[:5])
}
