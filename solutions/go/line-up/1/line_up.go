package lineup

import "fmt"

func Format(name string, number int) string {
	n:= number%100
    suffix:=""
    switch n{
        case 11,12,13:
        	suffix+="th"
        default:
        	    m:= n%10
    			switch m{
        			case 1:
						suffix+="st"
        			case 2:
						suffix+="nd"
        			case 3 :
						suffix+="rd"
        			default:
        				suffix+="th"
    			}
    }

    return fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, suffix)
}
