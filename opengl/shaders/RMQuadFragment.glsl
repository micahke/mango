#version 330 core

out vec4 FragColor;

in vec4 VertexColor;

void main() {
  FragColor = VertexColor;
  // FragColor = vec4(1.0, 0.0, 0.0, 1.0);
}

