package com.micahelias.graphic;

import com.micahelias.components.ModelComponent;
import com.micahelias.opengl.IndexBuffer;
import com.micahelias.opengl.Shader;
import com.micahelias.opengl.VertexArray;
import com.micahelias.opengl.VertexBuffer;
import com.micahelias.opengl.VertexBufferLayout;

public class Rect2D extends ModelComponent {


  public float width;
  public float height;


  float[] positions = {
    -0.5f, -0.5f,
    -0.5f, 0.5f,
    0.5f, 0.5f,
    0.5f, -0.5f
  };

  int[] indeces = {
    0, 1, 2,
    2, 3, 0
  };


  public Rect2D(float width, float height) {
    super();
    this.width = width;
    this.height = height;
  }


  public Rect2D(float x, float y, float width, float height) {
    super();

    entity.transform.position.x = x;
    entity.transform.position.y = y;

    this.width = width;
    this.height = height;

  }

  public void init() {
    initPositions();

    vao = new VertexArray();
    vbo = new VertexBuffer(positions);
    layout = new VertexBufferLayout();
    layout.pushf(2);
    vao.addBuffer(vbo, layout);

    ibo = new IndexBuffer(indeces);

    shader = new Shader("vertex2D.glsl", "fragment2D.glsl");
  }

  public void update() {};

  void initPositions() {
    this.positions = new float[] {
      0, 0,
      0, height,
      width, height,
      width, 0
    };
  }

}
