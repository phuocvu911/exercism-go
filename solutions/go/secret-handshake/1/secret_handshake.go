package secrethandshake
import "slices"
var number = []uint{1,2,4,8,16}
var action = []string{"wink", "double blink", "close your eyes", "jump", "reverse"}
func Handshake(code uint) []string {
	res:= []string{}
    for i, n:= range number{
        if code&n != 0{
            res = append(res, action[i])
        }
    }
    if len(res) > 0 && res[len(res)-1] == "reverse"{
        res = res[:len(res)-1]
    	slices.Reverse(res)
    }
    return res    
}