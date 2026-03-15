package letter

type FreqMap map[rune]int

func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

func ConcurrentFrequency(texts []string) FreqMap {
	res:= FreqMap{}
    ch := make(chan FreqMap)
    for _,text := range texts{
        go func(text string){
            ch <- Frequency(text)
        }(text)
    }
    for range texts{
        for k,v:= range <- ch{
            res[k] += v
        }
    }
    return res
}
