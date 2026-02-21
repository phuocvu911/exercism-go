package protein
import "errors"
var ErrStop = errors.New("Stop") // to be legal in package level
var ErrInvalidBase = errors.New("InvalidBase")
func FromRNA(rna string) ([]string, error) {
	res:= []string{}
    for i:=0; i<len(rna)-2; i+=3{
        acid, err := FromCodon(rna[i:i+3])
        if err == ErrStop && len(res)>=1{
            return res, nil
        } else if err == ErrInvalidBase{
            return nil, ErrInvalidBase
        }
        res = append(res, acid)
    }
    return res, nil
}

func FromCodon(codon string) (string, error) {
	
    switch codon{
        case "AUG":
        	return "Methionine", nil
        case "UUU", "UUC":
        	return "Phenylalanine", nil
        case "UUA", "UUG":
        	return "Leucine", nil
        case "UCU", "UCC", "UCA", "UCG":
        	return "Serine", nil
        case "UAU", "UAC":
        	return "Tyrosine", nil
        case "UGU", "UGC":
        	return "Cysteine", nil
        case "UGG":
        	return "Tryptophan", nil
        case "UAA", "UAG", "UGA":
        	return "", ErrStop
        default:
        	return "", ErrInvalidBase
    }
}
