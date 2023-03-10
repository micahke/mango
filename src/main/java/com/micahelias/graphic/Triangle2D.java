package com.micahelias.graphic;

import org.joml.Vector2f;

import com.micahelias.components.ModelComponent;
import com.micahelias.opengl.IndexBuffer;
import com.micahelias.opengl.Shader;
import com.micahelias.opengl.VertexArray;
import com.micahelias.opengl.VertexBuffer;
import com.micahelias.opengl.VertexBufferLayout;

public class Triangle2D extends ModelComponent {

  private Vector2f v1;
  private Vector2f v2;
  private Vector2f v3;


  float[] positions = {
    -0.5f, -0.5f,
    0.0f, 0.5f,
    0.5f, -0.5f
  };

  int[] indeces = {
    0, 1, 2
  };

  public Triangle2D(Vector2f xy1, Vector2f xy2, Vector2f xy3) {
    super();

    this.v1 = xy1;
    this.v2 = xy2;
    this.v3 = xy3;


    initVertices();
  }

  void initVertices() {
    this.positions = new float[]{
      v1.get(0), v1.get(1),
      v2.get(0), v2.get(1),
      v3.get(0), v3.get(1)
    };
  }

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
