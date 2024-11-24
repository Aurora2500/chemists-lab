package game

type System struct {
	PeriodicTable PeriodicTable
	CompoundTable CompoundTable
	Compounds     CompoundCollection
}

func NewSystem(compounds []Compound) *System {
	return &System{
		PeriodicTable: NewPeriodicTable(),
		CompoundTable: NewCompoundTable(),
		Compounds:     NewCollection(compounds),
	}
}

func (s *System) Bind() {
	s.PeriodicTable.Ssbo.BindShader(0)
	s.CompoundTable.Ssbo.BindShader(1)
	s.Compounds.Ssbo.BindShader(2)
}
