package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](slice []T, predicate func(T) bool) []T{
    res:= []T{}
    for _, v:= range slice{
        if predicate(v){
            res = append(res, v)
        }
    }
    return res
}

func Discard[T any](slice []T, predicate func(T) bool) []T{
    res:= []T{}
    for _, v:= range slice{
        if !predicate(v){
            res = append(res, v)
        }
    }
    return res
}
// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
