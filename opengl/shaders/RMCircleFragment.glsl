#version 330 core

out vec4 FragColor;

in vec4 VertexColor;
in vec2 TexCoord;

void main() {
  float distance = length(TexCoord - vec2(0.5, 0.5));

  if (distance <= 0.5) {
    FragColor = VertexColor;
  } else {
    discard;
  }
}

