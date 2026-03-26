package grep
import (
    "os"
    "bufio"
    "strings"
    "strconv"
)
// -i,-v,-x match condition
// -n,-l output style
func Search(pattern string, flags, files []string) []string {
	flagz := make(map[string]bool)
    for _,v := range flags{
        flagz[v] = true
    }
    if flagz["-i"]{
        pattern = strings.ToLower(pattern)
    }
    match := func(line string) bool{
        target := line //preserve input
        if flagz["-i"]{
        	target = strings.ToLower(target)
    	}
        var ok bool
        if flagz["-x"]{
            ok = target == pattern
        } else {
            ok = strings.Contains(target, pattern)
        }
        if flagz["-v"] {
            ok = !ok
        }
        return ok
    }
	isFiles := len(files) > 1
    res:= []string{}
    for _, filename:= range files{
        file, err:= os.Open(filename)
        if err != nil{
            continue
        }
        defer file.Close()
        scanner:= bufio.NewScanner(file)
        lineNum:=0
        for scanner.Scan(){
            lineNum++
            line := scanner.Text()
            if !match(line){
                continue
            }
            //if match
            if flagz["-l"]{
                res= append(res, filename)
                break
            }
            var b strings.Builder
            if isFiles{
                b.WriteString(filename)
                b.WriteByte(':')
            }
            if flagz["-n"]{
                b.WriteString(strconv.Itoa(lineNum))
                b.WriteByte(':')
            }
            b.WriteString(line)
            res= append(res, b.String())
        }
    }
    return res
}
