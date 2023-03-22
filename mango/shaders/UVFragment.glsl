#version 330 core

out vec4 FragColor;

in vec2 TexCoord;
uniform sampler2D uTexture;

uniform vec4 whiteChannel;
uniform vec4 blackChannel;

void main() {

  vec4 color = texture(uTexture, TexCoord);

  if (color.x > 0.5) {
    FragColor = whiteChannel;
  } else if (color.x < 0.5 && color.a != 0.0) {
    FragColor = blackChannel;
  } else {
    FragColor = color;
  }

}

