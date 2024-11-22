//shader vertex
#version 430

in vec3 pos;

uniform mat4 vp;

void main() {
	gl_Position = vp * vec4(pos, 1);
}

//shader fragment
#version 430

out vec4 col;

void main() {
	col = vec4(1, 0, 0, 1);
}