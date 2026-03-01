package rotationalcipher

func RotationalCipher(s string, shift int) string {
	res := ""
    for _,v:= range s{
        if v>='a' && v<='z'{
            res += string((v-'a' + rune(shift))%26 +'a')
    	} else if v>='A' && v<='Z'{
            res += string((v-'A' + rune(shift))%26 +'A')
    	} else{
            res+= string(v)
        }
    }
    return res
}
