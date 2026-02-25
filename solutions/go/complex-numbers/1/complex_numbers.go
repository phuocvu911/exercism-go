package complexnumbers
import "math"
// Define the Number type here.
type Number struct{
    Rea float64
    Imag float64 //imaginary part
}

func (n Number) Real() float64 {
	return n.Rea
}

func (n Number) Imaginary() float64 {
	return n.Imag
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.Rea+n2.Rea, n1.Imag + n2.Imag}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.Rea-n2.Rea, n1.Imag - n2.Imag}
}

func (n1 Number) Multiply(n2 Number) Number {    
	return Number{n1.Rea*n2.Rea-n1.Imag*n2.Imag, n1.Imag*n2.Rea + n1.Rea*n2.Imag}
}

func (n Number) Times(x float64) Number {
	n.Rea*=x
    n.Imag*=x
    return n
}

func (n1 Number) Divide(n2 Number) Number {
	a,b,c,d:= n1.Rea, n1.Imag, n2.Rea, n2.Imag
    n1.Rea = (a * c + b * d) / (c*c + d*d)
    n1.Imag = (b * c - a * d) / (c*c + d*d)
    return n1
}

func (n Number) Conjugate() Number {
	n.Imag = -n.Imag
    return n
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.Rea*n.Rea + n.Imag*n.Imag)
}

func (n Number) Exp() Number {
	a,b:= n.Rea, n.Imag
    n.Rea = math.Pow(math.E, a)*math.Cos(b)
    n.Imag = math.Pow(math.E, a)*math.Sin(b)
    return n
}
