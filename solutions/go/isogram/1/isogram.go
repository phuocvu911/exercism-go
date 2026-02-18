package isogram

import "strings"
func IsIsogram(w string) bool {
    word:= strings.ToLower(w)
	if word == ""{
        return true
    }
    for i:=0; i<len(word); i++{
        for j:=i+1; j<len(word); j++{
            if word[i] == word[j] && word[i] != '-' && word[i] != ' '{
                return false
            }
        }
    }
    return true
}
