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
    var res = []byte{byte(n&0x7F)}
    n>>=7 
    for n>0{
        res = append([]byte{byte(n&0x7F|0x80)}, res...)
        n>>=7
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
		if bitsAccumulated >= 32 {
			return nil, errors.New("overflow: value exceeds 32-bit unsigned integer")
		}
		current = (current << 7) | uint32(b&0x7f)
		bitsAccumulated += 7
 
		if b&0x80 == 0 { // we can check because in encode we already force the 8th bit to be 1 in every byte except the last one.
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