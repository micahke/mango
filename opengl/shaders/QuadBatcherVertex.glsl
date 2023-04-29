#version 330 core

layout(location = 0) in vec2 position;
layout(location = 1) in vec4 color;

uniform mat4 projection;
uniform mat4 view;
uniform mat4 model;

out vec4 Color;

void main() {

  gl_Position = projection * view * model * vec4(position.xy, 0.0, 1.0);
  Color = color;
}
