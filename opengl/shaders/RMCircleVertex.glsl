#version 330

// STANDARD
layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;
layout(location = 2) in vec2 texCoord;

// STANDARD
uniform mat4 u_model;
uniform mat4 u_view;
uniform mat4 u_projection;

out vec4 VertexColor;
out vec2 TexCoord;

void main()
{
    gl_Position = u_projection * u_view * u_model * position;
    VertexColor = color;
    TexCoord = texCoord;
}
