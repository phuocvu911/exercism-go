package allergies

func Allergies(allergies uint) []string {
	h:=map[uint]string{
        1: "eggs",
        2: "peanuts",
        4: "shellfish",
        8: "strawberries",
        16: "tomatoes",
        32: "chocolate",
        64: "pollen",
        128: "cats",
    }
    res:= []string{}
    for key, val := range h{
        if key&allergies != 0{ //bitwise AND 
            res = append(res, val)
        }
    }
    return res
}

func AllergicTo(allergies uint, allergen string) bool {
    for _, v := range Allergies(allergies){
        if allergen == v{
            return true
        }
    }
    return false
}