//shader vertex
#version 430 core

in vec2 pos;
in vec2 uv;

out vec2 texCoords;

void main() {
	gl_Position = vec4(pos, 0, 1);
	texCoords = uv;
}


//shader fragment
#version 430 core

uniform sampler2D tex;

in vec2 texCoords;

out vec4 col;

void main() {
	col = texture(tex, texCoords);
}