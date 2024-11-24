//shader vertex
#version 430

in vec3 pos;

layout (std430, binding = 0) buffer Positions {
	vec4 positions[];
};

uniform mat4 proj;
uniform mat4 view;

out vec3 normal;

void main() {
	gl_Position = proj * view * vec4(pos + vec3(positions[gl_InstanceID]), 1);
	normal = pos;
}

//shader fragment
#version 430

in vec3 normal;

out vec4 col;

const vec3 light = vec3(0.3, 0.9, 0.3);

void main() {
	float d = max(dot(normal, light), 0);
	float i = mix(0.3, 0.9, d);
	vec3 c = i * vec3(1, 0, 0);
	col = vec4(c, 1);
}