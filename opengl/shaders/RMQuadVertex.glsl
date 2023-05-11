#version 330

layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;

out vec4 VertexColor;

uniform mat4 u_model;
uniform mat4 u_view;
uniform mat4 u_projection;

void main()
{
    gl_Position = u_projection * u_view * u_model * position;
    VertexColor = color;
}
