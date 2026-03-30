package railfencecipher
var down = []int{1,1}
var up = []int{-1,1}

func Encode(m string, rails int) string {
	res := make([][]byte, rails)
    for i:= range res{
        res[i] = make([]byte, len(m))
    }
    a,b,v:=0,0,0
    encode := ""
    res[a][b] = m[v]
    for v < len(m){
        for range rails-1{
            v++
            if v>= len(m){
                break
            }
            a+=down[0]
            b+= down[1]
            res[a][b] = m[v]            
        }
        for range rails-1{
            v++
            if v>=len(m){
                break
            }
            a+=up[0]
            b+=up[1]
            res[a][b] = m[v]            
        }
    }
    for i:= range res{
        for j:= range res[i]{
            if res[i][j] != byte(0){
                encode += string(res[i][j])
            }
        }
    }
    return encode    
}

func Decode(m string, rails int) string {
	res := make([][]byte, rails)
    for i:= range res{
        res[i] = make([]byte, len(m))
    }
    a,b,v:=0,0,0
    posX:= make([]int, len(m))
    posY:= make([]int, len(m))
    res[a][b] = '!'
    for v < len(m){
        for range rails-1{
            v++
            if v>= len(m){
                break
            }
            a+=down[0]
            b+= down[1]
            posX[b] = a
            posY[b] = b
            res[a][b] = '!'            
        }
        for range rails-1{
            v++
            if v>=len(m){
                break
            }
            a+=up[0]
            b+=up[1]
            posX[b] = a
            posY[b] = b
            res[a][b] = '!'            
        }
    }
    x:=0
    for i:= range res{
        for j:= range res[i]{
            if res[i][j] == '!'{
                res[i][j] = m[x]
                x++
            }
        }
    }
    decode := string(m[0])
    for i:= 1; i<len(m); i++{
        decode += string(res[posX[i]][posY[i]])
    }
    return decode
}
