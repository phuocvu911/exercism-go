package sublist

func Sublist(l1, l2 []int) Relation {
	switch{
        case isEqual(l1,l2):
        	return RelationEqual
        case isSublist(l1,l2):
        	return RelationSublist
        case isSublist(l2,l1):
        	return RelationSuperlist
        default:
        	return RelationUnequal
    }
}

func isSublist(l1, l2 []int) bool{
    if len(l1) == 0{
        return true
    }
    if len(l1) > len(l2){
        return false
    }
    for i:= 0; i<=len(l2)-len(l1); i++{
        match := true
        for j:= 0; j<len(l1); j++{
            if l1[j] != l2[j+i]{
                match = false
                break
            } 
        }
        if match{
            return true
        }
    }
    return false
}

func isEqual(l1,l2 []int) bool{
    if len(l1) != len(l2){
        return false
    }
    for i:= 0; i<len(l1); i++{
        if l1[i] != l2[i]{
            return false
        }
    }
    return true
}