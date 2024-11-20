package rendering

type Mesh interface {
	Draw()
}

type ObjectMesh struct {
}

type IndexedMesh struct {
	ObjectMesh
}

type InstancedMesh struct {
	IndexedMesh
}
