package house
import "fmt"
var subjects = []string{
    "","",
    "the malt",
    "the rat",
    "the cat",
    "the dog",
    "the cow with the crumpled horn",
    "the maiden all forlorn",
    "the man all tattered and torn",
    "the priest all shaven and shorn",
    "the rooster that crowed in the morn",
    "the farmer sowing his corn",
    "the horse and the hound and the horn",
}

var verbs = []string{
    "","",
    "lay in",
    "ate",
    "killed",
    "worried",
    "tossed",
    "milked",
    "kissed",
    "married",
    "woke",
    "kept",
    "belonged to",
}
func Verse(v int) string {
    res:="This is"
    for v>1{
        res+=fmt.Sprintf(" %s\nthat %s", subjects[v], verbs[v])
        v--
    }
    res+= " the house that Jack built."
    return res
}

func Song() string {
	res:=""
    for v:=1; v<=12; v++{
        res+= Verse(v)
        if v<12{
            res+= "\n\n"
        }
    }
    return res
}
