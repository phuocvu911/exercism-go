package parsinglogfiles
import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<[~*=-]*\>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count:= 0
    re:= regexp.MustCompile(`(?i)"[^"]*password"`)
    for _, line:= range lines{
        if re.MatchString(line){
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re:= regexp.MustCompile(`end-of-line\d*`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re:= regexp.MustCompile(`User\s+(\S+)`) //() for capturing the username that we can access later by m[1], m[0] is full match
    for i, line := range lines{
        m:= re.FindStringSubmatch(line)
        if m != nil{
            lines[i] = "[USR] " + m[1] + " " +line 
    	}
    }
    return lines
}
