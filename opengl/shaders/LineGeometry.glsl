#version 330

layout(lines) in;
layout(triangle_strip, max_vertices = 4) out;

uniform float thickness;

void main() {
    vec3 p1 = gl_in[0].gl_Position.xyz;
    vec3 p2 = gl_in[1].gl_Position.xyz;
    vec3 dir = normalize(p2 - p1);
    vec3 normal = vec3(-dir.y, dir.x, 0.0);
    vec4 offset = vec4(normal * thickness, 0.0);
    gl_Position = gl_in[0].gl_Position + offset;
    EmitVertex();
    gl_Position = gl_in[1].gl_Position + offset;
    EmitVertex();
    gl_Position = gl_in[0].gl_Position - offset;
    EmitVertex();
    gl_Position = gl_in[1].gl_Position - offset;
    EmitVertex();
    EndPrimitive();
}
