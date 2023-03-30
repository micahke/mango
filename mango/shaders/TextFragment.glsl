#version 330 core

out vec4 FragColor;

in vec2 TexCoord;
uniform sampler2D uTexture;

void main() {

  vec4 clr = texture(uTexture, TexCoord);

  if (clr.a == 0.0){
    FragColor = vec4(1.0, 0.0, 0.0, 1.0);
  } else {
    // FragColor = clr;
    FragColor = vec4(1.0, 1.0, 1.0, clr.x);
  }

  // FragColor = clr;
  // // FragColor = vec4(1.0, 0.0, 0.0, 1.0);

}

