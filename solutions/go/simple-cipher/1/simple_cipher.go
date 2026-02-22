package cipher
import "strings"

type shift struct{
    Distance int
}

type vigenere struct{
    Key []int
}
// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
    if distance == 0 || distance > 25 || distance < -25{
        return nil
    }
	distance = (distance +26)%26 //dealing with negative
    return shift{distance}
}

func (c shift) Encode(input string) string {	
    res := ""
    for _,v := range strings.ToLower(input){
        if v>= 'a' && v<= 'z'{
            res += string((v-'a'+rune(c.Distance))%26+'a')
        }
    }
    return res
}

func (c shift) Decode(input string) string {
    res := ""
    for _,v := range strings.ToLower(input){
        if v>= 'a' && v<= 'z'{
            res += string((v-'a'-rune(c.Distance)+26)%26+'a')
        }
    }
    return res	
}

func NewVigenere(key string) Cipher {
	res := []int{}
    if key == ""{
        return nil
    }
	//Values consisting entirely of the letter 'a' are disallowed
    f:= func(r rune) bool{
        return r>'a' && r<='z'
    }
    f1:= func(r rune) bool{
        return r<'a' || r>'z'
    }
    if !strings.ContainsFunc(key, f) || strings.ContainsFunc(key, f1){
        return nil
    }
    for _, v:= range key{
        res= append(res, int(v-'a'))
    }
    return vigenere{res}
}

func (v vigenere) Encode(input string) string {
	indexKey := 0
    res := ""
    for _, ch:= range strings.ToLower(input){
        shift:= v.Key[indexKey%len(v.Key)] // so that the key loops over
        if ch>= 'a' && ch<= 'z'{
            res += string((ch-'a'+rune(shift))%26+'a')
        } else {
            continue
        }
        indexKey++
    }
    return res
}

func (v vigenere) Decode(input string) string {
	indexKey := 0
    res := ""
    for _, ch:= range strings.ToLower(input){
        shift:= v.Key[indexKey%len(v.Key)] // so that the key loops over
        if ch>= 'a' && ch<= 'z'{
            res += string((ch-'a'-rune(shift)+26)%26+'a')
        } else {
            continue
        }
        indexKey++
    }
    return res
}