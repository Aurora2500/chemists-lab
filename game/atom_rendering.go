package game

import "chemists-lab/rendering"

type AtomInfo struct {
	colorSize rendering.Vec4
}

type PeriodicTable struct {
	Ssbo *rendering.Ssbo[AtomInfo]
	data []AtomInfo
}

func NewPeriodicTable() PeriodicTable {
	table := PeriodicTable{
		Ssbo: rendering.NewSsbo[AtomInfo](),
	}

	table.data = append(table.data,
		AtomInfo{colorSize: rendering.Vec4{1, 1, 1, 0.6}},
		AtomInfo{colorSize: rendering.Vec4{1, .1, .1, 0.9}},
		AtomInfo{colorSize: rendering.Vec4{1, 1, .1, 1.1}},
		AtomInfo{colorSize: rendering.Vec4{.3, .3, .3, 1}},
	)

	table.Ssbo.Allocate(len(table.data), rendering.DYNAMIC_DRAW)
	table.Ssbo.Update(table.data)

	return table
}
