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

type CompoundTable struct {
	Ssbo  *rendering.Ssbo[CompoundInfo]
	Table []CompoundInfo
}

func NewCompoundTable() CompoundTable {
	type Vec3 = rendering.Vec3

	table := []CompoundInfo{
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
	ssbo.Allocate(len(table), rendering.STATIC_DRAW)
	ssbo.Update(table)

	return CompoundTable{
		Ssbo:  ssbo,
		Table: table,
	}
}

type Compound struct {
	Rotation Quat
	Pos      Vec3
	Compound int32
}

type CompoundCollection struct {
	Ssbo      *rendering.Ssbo[Compound]
	Compounds []Compound
}

func NewCollection(collection []Compound) CompoundCollection {
	ssbo := rendering.NewSsbo[Compound]()
	ssbo.Allocate(len(collection), rendering.DYNAMIC_DRAW)
	ssbo.Update(collection)
	return CompoundCollection{
		Ssbo:      ssbo,
		Compounds: collection,
	}
}

func (col *CompoundCollection) NumCompounds() int {
	return len(col.Compounds)
}

func (col *CompoundCollection) Update(update_func func(*Compound)) {
	for i := range col.Compounds {
		update_func(&col.Compounds[i])
	}
	col.Ssbo.Update(col.Compounds)
}
