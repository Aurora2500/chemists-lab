//shader vertex
#version 410

in vec2 pos;
in vec3 col;

out vec3 color;

void main() {
	gl_Position = vec4(pos, 0., 1.);
	color = col;
}

//shader fragment
#version 410

in vec3 color;

out vec4 c;

void main() {
	c = vec4(color, 1.);
}