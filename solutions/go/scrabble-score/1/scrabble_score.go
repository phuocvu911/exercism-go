package scrabble
import "strings"
func Score(word string) int {
	w := strings.ToUpper(word)
    p:= 0
    for _, v := range w{
        switch v{
            case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
            	p+=1
            case 'D', 'G':
            	p+=2
            case 'B', 'C', 'M', 'P':
            	p+=3
            case 'F', 'H', 'V', 'W', 'Y':
            	p+=4
            case 'K':
            	p+=5
            case 'J', 'X':
            	p+=8
            case 'Q', 'Z':
            	p+=10
        }
    }
    return p
}
