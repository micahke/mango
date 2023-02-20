package com.micahelias.graphic;

import com.micahelias.components.ModelComponent;
import com.micahelias.opengl.IndexBuffer;
import com.micahelias.opengl.Shader;
import com.micahelias.opengl.VertexArray;
import com.micahelias.opengl.VertexBuffer;
import com.micahelias.opengl.VertexBufferLayout;

public class Triangle2D extends ModelComponent {


  float[] positions = {
    -0.5f, -0.5f,
    0.0f, 0.5f,
    0.5f, -0.5f
  };

  int[] indeces = {
    0, 1, 2
  };

  public void init() {
    vao = new VertexArray();
    vbo = new VertexBuffer(positions);
    layout = new VertexBufferLayout();
    layout.pushf(2);
    vao.addBuffer(vbo, layout);

    ibo = new IndexBuffer(indeces);

    shader = new Shader("vertex2D.glsl", "fragment2D.glsl");
  };

  public void update() {};

}
