#version 330

layout (location = 0) in vec4 position;
layout (location = 1) in vec2 textureCoord;

out vec2 TexCoord;

void main()
{
  gl_Position = position;
  TexCoord = textureCoord;
}
