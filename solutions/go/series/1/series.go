package series

func All(n int, s string) []string {
	res := []string{}
    for i:= 0; i<=len(s)-n; i++{
        res = append(res, s[i:i+n])
    }
    return res
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
