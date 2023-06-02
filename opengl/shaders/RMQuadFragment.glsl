#version 330 core

out vec4 FragColor;

in vec4 VertexColor;
in vec2 TexCoord;

uniform sampler2D uTexture;
uniform int isTextured;

void main() {
  if (isTextured == 0) {
    FragColor = texture(uTexture, TexCoord);
  } else {
    FragColor = VertexColor;
  }
  // FragColor = vec4(1.0, 0.0, 0.0, 1.0);
    // FragColor = texture(uTexture, TexCoord);
}

