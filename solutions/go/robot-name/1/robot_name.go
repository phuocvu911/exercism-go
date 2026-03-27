package robotname
import (
    "math/rand"
    "fmt"
    "errors"
)
// Define the Robot type here.
type Robot struct{
    name string
}

var (
    letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    used = make(map[string]bool)
)

func generateName() string{
    return fmt.Sprintf("%c%c%03d",
                  letter[rand.Intn(26)],
                  letter[rand.Intn(26)],
                   rand.Intn(1000))
}

func (r *Robot) Name() (string, error) {
	if r.name != ""{
        return r.name, nil
    }
    if len(used) == maxNames{
        return "", errors.New("no more unique name")
    }
    for { //try new name until its unique
        name:= generateName()
        if used[name]{
            continue
        }
        r.name = name
        used[name] = true
        return r.name, nil
    }
}

func (r *Robot) Reset() {
	if r.name != ""{
        used[r.name] = false
        r.name  = ""
    }
}