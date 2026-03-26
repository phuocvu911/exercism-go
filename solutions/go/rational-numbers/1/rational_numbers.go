package rationalnumbers
import "math"
type Rational struct {
	n,d int
}

// Reduce simplifies a Rational, eg changing Rational{4, 2} into Rational{2, 1}.
func abs(a int) int{
    if a<0{
        return -a
    }
    return a
}
func Gcd(a,b int) int{
	for b!=0{
        a,b = b, a%b
    }
    return a
}
func (r Rational) Reduce() Rational {
	gcd:= abs(Gcd(r.n,r.d))
    if r.d < 0{
        gcd *=-1
    }
    r.n /= gcd
    r.d /= gcd
    return r
}

func (r Rational) Add(s Rational) Rational {
	r.n = r.n*s.d + r.d*s.n
    r.d *= s.d
    return r.Reduce()
}

func (r Rational) Sub(s Rational) Rational {
	r.n = r.n*s.d - r.d*s.n
    r.d *= s.d
    return r.Reduce()
}

func (r Rational) Mul(s Rational) Rational {
	r.n *= s.n
	r.d *= s.d
    return r.Reduce()
}

func (r Rational) Div(s Rational) Rational {
	r.n *= s.d
	r.d *= s.n
    return r.Reduce()
}

func (r Rational) Abs() Rational {
	if r.n <0{
        r.n *=-1
    }
    if r.d <0{
        r.d*=-1
    }
    return r.Reduce()
}

// Compute r ^ power, a rational raised to an int exponent.
func (r Rational) Exprational(power int) Rational {
	if power == 0{
        return Rational{1,1}
    }
    if power <0{
        r.n, r.d = r.d, r.n
        power = -power
    }
    a,b := r.n, r.d
    for i:=1; i<power;i++{
        r.n*=a
        r.d*=b
    }
    return r.Reduce()
}

// Compute base ^ r, an int raised to a rational.
func (r Rational) Expreal(base int) float64 {
	return math.Pow(float64(base), float64(r.n)/float64(r.d))
}
