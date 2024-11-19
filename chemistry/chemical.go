package chemistry

type chemicalAtom struct {
	atom  *Atom
	bonds []chemicalAtom
}

type Chemical struct {
	atoms []Atom
}
