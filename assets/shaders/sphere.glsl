//shader vertex
#version 430

in vec3 pos;

uniform mat4 proj;
uniform mat4 view;
uniform mat4 model;

void main() {
	gl_Position = proj * view * model * vec4(pos, 1);
}

//shader fragment
#version 430

out vec4 col;

void main() {
	col = vec4(1, 0, 0, 1);
}