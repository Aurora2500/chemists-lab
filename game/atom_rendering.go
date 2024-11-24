package game

import "chemists-lab/rendering"

type AtomInfo struct {
	colorSize rendering.Vec4
}

type PeriodicTable struct {
	Ssbo *rendering.Ssbo[AtomInfo]
	data []AtomInfo
}

func NewPeriodicTable() *PeriodicTable {
	table := PeriodicTable{
		Ssbo: rendering.NewSsbo[AtomInfo](),
	}

	table.data = append(table.data,
		AtomInfo{colorSize: rendering.Vec4{1, 1, 1, 0.4}},
		AtomInfo{colorSize: rendering.Vec4{1, .1, .1, 0.9}},
	)

	table.Ssbo.Allocate(len(table.data), rendering.DYNAMIC_DRAW)
	table.Ssbo.Update(table.data)

	return &table
}
