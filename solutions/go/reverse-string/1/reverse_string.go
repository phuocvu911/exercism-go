package reverse

func Reverse(input string) string {
	res:= ""
    for _,v:= range input{
        res = string(v) + res
    }
    return res
}
