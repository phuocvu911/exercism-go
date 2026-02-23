package triangle

type Kind string

const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

func KindFromSides(a, b, c float64) Kind {
    if a<=0 || b<=0 || c<=0{
        return NaT
    }
    if a+b<c || b+c<a || a+c<b{
        return NaT
    } 
    if a ==b && b==c{
        return Equ
    }
    if a == b || b==c || a==c{
        return Iso
    }
    return Sca
}