package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	res:= ""
    for _,v:= range log{
        switch  v{
            case '❗':
            res = "recommendation"
            goto End
            case '🔍':
            res = "search"
            goto End
            case '☀':
            res = "weather"
			goto End        
        }
    }
    if res == ""{
        res = "default"
    }
    End:
    return res
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
    for i, c:= range runes{
        if c == oldRune{
            runes[i] = newRune
        }
    }
    return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
