package flattenarray

func Flatten(nested any) []any {
	res:= []any{}
    data, ok:= nested.([]any) //check if input is a slice
    if !ok{
        return res
    } else{
        for _, v:= range data{
            if v == nil{
                continue
            }
            if inner, ok1:= v.([]any); ok1{
                res = append(res, Flatten(inner)...) //have to catch those deep layer from recursive and add it to res.
            } else{
                res= append(res, v)
            }
        }
    }
    return res
}
