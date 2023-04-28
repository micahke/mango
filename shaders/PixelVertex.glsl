#version 330 core

layout(location = 0) in vec4 position;

uniform mat4 projection;
uniform mat4 view;
uniform mat4 model;

uniform float uPixelSize;

void main() {

  gl_Position = projection * view * model * position;
  gl_PointSize = uPixelSize;
}
