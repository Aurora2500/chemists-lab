//shader vertex
#version 430

in vec3 pos;

layout(std430, binding = 0) buffer PeriodicTable {
	vec4 atomInfo[];
};

struct Atom {
	vec3 position;
	int atomicNumber;
};

layout (std430, binding = 1) buffer Positions {
	Atom atoms[];
};

uniform mat4 proj;
uniform mat4 view;

out vec3 normal;
out vec3 col;

void main() {
	Atom atom = atoms[gl_InstanceID];
	col = atomInfo[atom.atomicNumber].rgb;
	float size = atomInfo[atom.atomicNumber].w;
	vec3 position = (size * pos) + atom.position;

	gl_Position = proj * view * vec4(position, 1);
	normal = pos;
}

//shader fragment
#version 430

in vec3 normal;
in vec3 col;

out vec4 frag_col;

const vec3 light = vec3(0.3, 0.9, 0.3);

void main() {
	float d = max(dot(normal, light), 0);
	float i = mix(0.3, 0.9, d);
	frag_col = vec4(col * i, 1);
}