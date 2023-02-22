#version 330 core

layout(location = 0) in vec4 position;

uniform mat4 projectionMatrix;
uniform mat4 modelMatrix;


void main() {
  gl_Position = projectionMatrix * modelMatrix * vec4(position.xy, 0.0, 1.0);
}
