package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int{
    if time == 0 {
        time +=2
    }
    return len(layers) * time
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    no := 0
    sau := 0.0
    for _,i := range layers {
        if i =="noodles" {
            no +=50
        } else if i == "sauce" {
            sau += 0.2
        }
    }
    return no, sau
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(l1,l2 []string) {
    l2[len(l2)-1] = l1[len(l1)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amount []float64, portion int) []float64 {
    res := make([]float64,len(amount))
    for i,val := range amount {
        res[i] = val*float64(portion)/2.0
    }
    return res
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
