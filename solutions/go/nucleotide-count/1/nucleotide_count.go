package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A':0, 'C':0, 'G':0, 'T':0}
    for _,v := range d{
        if _, ok:=h[v]; !ok{
            return Histogram{}, errors.New("Not Valid")
        }
        h[v]++
    }
	return h, nil
}
