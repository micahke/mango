package com.micahelias.components;

import org.joml.Matrix4f;

import com.micahelias.opengl.IndexBuffer;
import com.micahelias.opengl.Shader;
import com.micahelias.opengl.VertexArray;
import com.micahelias.opengl.VertexBuffer;
import com.micahelias.opengl.VertexBufferLayout;
import com.micahelias.util.Color;

public class ModelComponent extends Component {

  public VertexArray vao;
  public VertexBuffer vbo;
  public VertexBufferLayout layout;
  public IndexBuffer ibo;
  public Shader shader;
  public Matrix4f modelMatrix;

  public Color color = new Color(0, 0, 0);


  // GETTERS

  public VertexArray getVAO() {
    return this.vao;
  }
  
  public VertexBuffer setVBO() {
    return this.vbo;
  }


  public VertexBufferLayout getLayout() {
    return this.layout;
  }

  public IndexBuffer getIBO() {
    return this.ibo;
  }

  public Shader getShader() {
    return this.shader;
  }

  public Matrix4f getModelMatrix() {
    return this.modelMatrix;
  }

  public Color getColor() {
    return this.color;
  }

  // SETTERS

  public void setVAO(VertexArray vao) {
    this.vao = vao;
  }

  
  public void setVBO(VertexBuffer vbo) {
    this.vbo = vbo;
  }

  
  public void setLayout(VertexBufferLayout layout) {
    this.layout = layout;
  }

  public void setIBO(IndexBuffer ibo) {
    this.ibo = ibo;
  }

  public void setShader(Shader shader) {
    this.shader = shader;
  }

  public void setModelMatrix(Matrix4f matrix) {
    this.modelMatrix = matrix;
  }

  public ModelComponent setColor(Color color) {
    this.color = color;
    return this;
  }

  public void init() {};
  public void update() {};


}
