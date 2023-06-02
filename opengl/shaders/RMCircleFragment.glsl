#version 330 core

out vec4 FragColor;

in vec4 VertexColor;
in vec2 TexCoord;

uniform int isTextured;
uniform sampler2D uTexture;

void main() {
  float distance = length(TexCoord - vec2(0.5, 0.5));

  if (distance <= 0.5) {
    if (isTextured == 0) {
      FragColor = texture(uTexture, TexCoord);
    } else {
    FragColor = VertexColor;
    }
  } else {
    discard;
  }
}

