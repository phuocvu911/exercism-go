package foodchain
import "fmt"
var animal = []string{"","fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var cmt = map[string]string{
    "spider": "It wriggled and jiggled and tickled inside her.\n",
    "bird": "How absurd to swallow a bird!\n",
    "cat": "Imagine that, to swallow a cat!\n",
    "dog": "What a hog, to swallow a dog!\n",
    "goat": "Just opened her throat and swallowed a goat!\n",
    "cow": "I don't know how she swallowed a cow!\n",
}

func Verse(v int) string {
	start := fmt.Sprintf("I know an old lady who swallowed a %s.\n", animal[v])
    if animal[v] == "horse"{
        return start + "She's dead, of course!"
    }
    end := fmt.Sprintf("I don't know why she swallowed the fly. Perhaps she'll die.")
    mid:= ""
    if v>1{
		comment:= cmt[animal[v]]
        catch:=""
        for v>1{
            if animal[v-1] == "spider"{
                catch+= fmt.Sprintf("She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.\n")
            } else{
            catch+= fmt.Sprintf("She swallowed the %s to catch the %s.\n", animal[v], animal[v-1])
            }
            v--
        }
        mid+=comment + catch    	
    }
    return start + mid + end
}

func Verses(start, end int) string {
	res:=""
    for i:= start; i<= end; i++{
        res+= Verse(i)
        if i< end{
            res += "\n\n"
        }
    }
    return res
}

func Song() string {
	return Verses(1,8)
}
