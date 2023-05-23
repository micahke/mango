#version 330 core

out vec4 FragColor;
uniform sampler2D uTexture;

in vec2 TexCoord;

void main()
{
  FragColor = texture(uTexture, TexCoord);
  // FragColor = vec4(1.0, 0.0, 0.0, 1.0);
}
