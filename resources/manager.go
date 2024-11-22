package resources

import (
	"chemists-lab/rendering"
	"os"
	"path/filepath"
)

type Manager struct {
	base    string
	shaders map[string]*rendering.Shader
}

func NewManager(base string) *Manager {
	return &Manager{
		base:    base,
		shaders: make(map[string]*rendering.Shader),
	}
}

func (mgr *Manager) GetShader(name string) *rendering.Shader {
	if shader, ok := mgr.shaders[name]; ok {
		return shader
	}

	//TODO: make sure shader exists
	shader_path := filepath.Join(mgr.base, "shaders", name+".glsl")
	src, err := os.ReadFile(shader_path)
	if err != nil {
		panic(err)
	}

	shader, err := rendering.NewShader(string(src))
	if err != nil {
		panic(err)
	}
	return shader
}
