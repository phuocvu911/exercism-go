package yacht

func Score(dice []int, category string) int {
	var count [7]int
    sum:=0
    for _,d:= range dice{
        count[d]++
        sum+=d
    }
    switch category{
        case "ones":
        	return count[1] * 1
		case "twos":
			return count[2] * 2
		case "threes":
			return count[3] * 3
		case "fours":
			return count[4] * 4
		case "fives":
			return count[5] * 5
		case "sixes":
			return count[6] * 6
        case "choice":
        	return sum
        case "full house":
        hasThree, hasTwo:= false, false
        for _,v:= range count{
            if v == 3{
                hasThree = true
            }
            if v == 2{
                hasTwo = true
            }
        }
        if hasThree && hasTwo{
            return sum
        }
        return 0
        case "four of a kind":
        for i, v:= range count{
            if v>=4{
                return 4*i
            }
        }
        return 0
        case "little straight":
        for i:=1; i<=5; i++{
            if count[i] != 1{
                return 0
            }
        }
        return 30
        case "big straight":
        for i:=2; i<=6; i++{
            if count[i] != 1{
                return 0
            }
        }
        return 30
        case "yacht":
        for _,v:= range count{
            if v == 5{
                return 50
            }
        }
        return 0
    }
    return 0
}
