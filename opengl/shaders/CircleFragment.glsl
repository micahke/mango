#version 330 core

out vec4 FragColor;

in vec2 TexCoord;

uniform vec4 uColor;

void main() {

  float distance = length(TexCoord - vec2(0.5, 0.5));

  if (distance <= 0.5) {
    FragColor = uColor;
  } else {
    FragColor = vec4(0.0, 0.0, 0.0, 0.0);
  }

}
