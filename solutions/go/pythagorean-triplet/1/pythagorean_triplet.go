package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	res:= []Triplet{}
    for i:= min; i<=max; i++{
        for j:=i+1; j<= max; j++{
            for k:= j+1; k<= max;k++{
                if i*i + j*j == k*k{
                    res = append(res, Triplet{i,j,k})
                }
            }
        }
    }
    return res
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	checker:= Range(1, p/2)
    res:= []Triplet{}
    for _, t:= range checker{
        if t[0]+t[1]+t[2] == p{
            res = append(res, t)
        }
    }
    return res
}
