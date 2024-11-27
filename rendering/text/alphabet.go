package text

type Alphabet interface {
	Runes() []rune
}

type RuneArray []rune

func (r RuneArray) Runes() []rune {
	return r
}

type RuneRange struct {
	From, To rune
}

func (rr RuneRange) Runes() []rune {
	var runes []rune
	for r := rr.From; r <= rr.To; r++ {
		runes = append(runes, r)
	}
	return runes
}
