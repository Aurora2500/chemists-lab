//shader vertex
#version 430

in vec3 pos;

out vec3 vPos;
out int compound;

void main() {
	compound = gl_InstanceID;
	vPos = pos;
}

//shader geometry
#version 430

layout(triangles) in;
layout(triangle_strip, max_vertices = 96) out;

layout(std430, binding = 0) buffer PeriodicTable {
	vec4 atomInfo[];
};

struct Atom {
	vec3 position;
	int atomicNumber;
};

struct CompoundInfo {
	Atom atoms[4];
	int numAtoms;
};

layout (std430, binding = 1) buffer CompoundTable {
	CompoundInfo compoundInfo[];
};

struct Compound {
	vec4 rotation;
	vec3 position;
	int compound;
};

layout (std430, binding = 2) buffer Compounds {
	Compound compounds[];
};

uniform mat4 proj;
uniform mat4 view;

in vec3 vPos[];
in int compound[];

out vec3 normal;
out vec3 col;

mat4 quat_pos_to_homo(vec4 quat, vec3 pos) {
	float x = quat.x;
	float y = quat.y;
	float z = quat.z;
	float s = quat.w;
	return mat4(
		1 - 2*y*y - 2*z*z, 2*x*y + 2*s*z, 2*x*z - 2*s*y, 0,
		2*x*y - 2*s*z, 1 - 2*x*x - 2*z*z, 2*y*z + 2*s*x, 0,
		2*x*z + 2*s*y, 2*y*z - 2*s*x, 1 - 2*x*x - 2*y*y, 0,
		pos.x, pos.y, pos.z, 1
	);
}

void main() {
	Compound currCompound = compounds[compound[0]];
	vec4 quat = currCompound.rotation.yzwx;
	vec3 pos = currCompound.position;
	mat4 model = quat_pos_to_homo(quat, pos);

	CompoundInfo cinfo = compoundInfo[currCompound.compound];

	for (int ai = 0; ai < cinfo.numAtoms; ai++) {
		Atom atom = cinfo.atoms[ai];
		vec4 ainfo = atomInfo[atom.atomicNumber];
		for (int i = 0; i < 3; i++) {
			vec3 p = ainfo.w * vPos[i] + atom.position;
			gl_Position = proj * view * model * vec4(p, 1);
			col = ainfo.rgb;
			normal = mat3(model) * vPos[i];
			EmitVertex();
		}
		EndPrimitive();
	}
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