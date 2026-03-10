package matchingbrackets
var pairs = map[rune]rune{')':'(', ']':'[', '}':'{'}
var openers = map[rune]bool{'(': true, '[': true, '{': true}
func Bracket(s string) bool {
	checker:= []rune{}
    for _, r:=  range s{
        if openers[r]{
            checker = append(checker, r)
        } else if v,ok:= pairs[r]; ok{ //if that rune is a closer, check if the checker have correspond opener
            if len(checker) == 0 || v != checker[len(checker)-1]{
                return false
            }
            checker = checker[:len(checker)-1] // if it has the correspond opener in the checker, remove that opener, finish checking that pair
        }
    }
    return len(checker) == 0
}
