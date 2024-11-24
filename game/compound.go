package game

import "chemists-lab/rendering"

type Atom struct {
	Position     rendering.Vec3
	AtomicNumber int32
}

type CompoundInfo struct {
	Atoms    [4]Atom
	NumAtoms int32
	_        [12]byte
}

func NewCompoundInfo() *rendering.Ssbo[CompoundInfo] {
	type Vec3 = rendering.Vec3

	cinfo := []CompoundInfo{
		{
			Atoms: [4]Atom{
				{Position: Vec3{-.4, 0, 0}, AtomicNumber: 1},
				{Position: Vec3{+.4, 0, 0}, AtomicNumber: 1},
			},
			NumAtoms: 2,
		},
		{
			Atoms: [4]Atom{
				{Position: Vec3{0, 0.2, 0}, AtomicNumber: 1},
				{Position: Vec3{-0.7, -0.2, 0}, AtomicNumber: 0},
				{Position: Vec3{+0.7, -0.2, 0}, AtomicNumber: 0},
			},
			NumAtoms: 3,
		},
	}

	ssbo := rendering.NewSsbo[CompoundInfo]()
	ssbo.Allocate(len(cinfo), rendering.STATIC_DRAW)
	ssbo.Update(cinfo)

	return ssbo
}
