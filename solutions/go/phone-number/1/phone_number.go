package phonenumber
import (
    "regexp"
    "strings"
    "errors"
    "fmt"
)
var Err = errors.New("invalid")
func Number(s string) (string, error) {
	s= strings.TrimPrefix(s, "1")
	s= strings.TrimPrefix(s, "+1")
    re:= regexp.MustCompile(`[2-9][0-9]+`)
    res:= re.FindAllString(s, -1)
    a:= strings.Join(res, "")
    if len(a) != 10 {
        return "", Err
    }
    return a, nil
}

func AreaCode(s string) (string, error) {
	res, err:= Number(s)
    if err!=nil{
        return "", err
    }
    return res[:3], nil
}

func Format(s string) (string, error) {
	res, err:= Number(s)
    if err!=nil{
        return "", err
    }
    return fmt.Sprintf("(%s) %s-%s", res[:3], res[3:6], res[6:]), nil
}
