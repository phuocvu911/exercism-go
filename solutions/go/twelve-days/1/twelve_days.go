package twelve
import "fmt"
func Verse(i int) string {
	items:= []string{" twelve Drummers Drumming,", " eleven Pipers Piping,", " ten Lords-a-Leaping,", " nine Ladies Dancing,", " eight Maids-a-Milking,", " seven Swans-a-Swimming,", " six Geese-a-Laying,", " five Gold Rings,", " four Calling Birds,", " three French Hens,", " two Turtle Doves,", " a Partridge in a Pear Tree."}
    days:= []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
    ver:= fmt.Sprintf("On the %s day of Christmas my true love gave to me:", days[i-1])
    if i == 1{
        ver+=items[12-i]
        return ver
    }
    for k:=12-i; k<11; k++{
        ver+= items[k]
    }
    ver += " and" + items[11]
    return ver
}

func Song() string {
	res :=""
    for i:= 1; i<=12; i++{
        res+= Verse(i)
        if i<12{
            res += "\n"
        }
    }
    return res
}
