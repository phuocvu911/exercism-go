package binarysearch
import "sort"
func SearchInts(list []int, key int) int {
	res:= sort.SearchInts(list, key)
    if res < len(list) && list[res] == key{
        return res
    }
    return -1
}
