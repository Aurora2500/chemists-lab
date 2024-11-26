package resources

import (
	"chemists-lab/rendering"
	"os"
	"path/filepath"

	"golang.org/x/image/font/opentype"
)

type Manager struct {
	base    string
	shaders map[string]*rendering.Shader
	fonts   map[string]*opentype.Font
}

func NewManager(base string) *Manager {
	return &Manager{
		base:    base,
		shaders: make(map[string]*rendering.Shader),
		fonts:   make(map[string]*opentype.Font),
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

func (mgr *Manager) GetFont(name string) *opentype.Font {
	if font, ok := mgr.fonts[name]; ok {
		return font
	}

	font_path := filepath.Join(mgr.base, "fonts", name+".ttf")
	data, err := os.ReadFile(font_path)
	if err != nil {
		panic(err)
	}
	font, err := opentype.Parse(data)
	if err != nil {
		panic(err)
	}
	mgr.fonts[name] = font
	return font
}
