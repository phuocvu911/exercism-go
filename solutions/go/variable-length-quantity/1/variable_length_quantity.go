package variablelengthquantity
import "errors"
func EncodeVarint(input []uint32) []byte {
	var res []byte
    for _, v:= range input{
        res = append(res, encode(v)...)
    }
    return res
}

func encode(n uint32) []byte {
    var res []byte
    res = append(res, byte(n&0x7F)) //find the least significant 7bits
    n>>=7 
    for n>0{
        res = append(res, byte(n&0x7F)|0x80)
        n>>=7
    }
    //reverse a slice without import, because we want the most significant bit
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
    	res[i], res[j] = res[j], res[i]
	}
    return res

}

func DecodeVarint(data []byte) ([]uint32, error) {
	var result []uint32
	var current uint32
	var bitsAccumulated uint
	inSequence := false
 
	for _, b := range data {
		inSequence = true
		// Check for overflow: we allow at most 32 bits total.
		// Each byte contributes 7 bits; after 5 bytes we'd have 35 bits.
		if bitsAccumulated >= 32 {
			return nil, errors.New("overflow: value exceeds 32-bit unsigned integer")
		}
		current = (current << 7) | uint32(b&0x7F)
		bitsAccumulated += 7
 
		if b&0x80 == 0 {
			// Last byte of this number, reset and go on next num
			result = append(result, current)
			current = 0
			bitsAccumulated = 0
			inSequence = false
		}
	}
 
	if inSequence {
		return nil, errors.New("incomplete sequence: input ended mid-value")
	}
 
	return result, nil
}

