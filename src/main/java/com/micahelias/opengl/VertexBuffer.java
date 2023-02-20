package com.micahelias.opengl;

import static org.lwjgl.opengl.GL15.*;

/**
 * VertexBuffer
 */
public class VertexBuffer {

  private int m_RendererID;
  
  public VertexBuffer(float[] data) {
    m_RendererID = glGenBuffers();
    glBindBuffer(GL_ARRAY_BUFFER, m_RendererID);
    glBufferData(GL_ARRAY_BUFFER, data, GL_STATIC_DRAW);
  }


  void bind() {
    glBindBuffer(GL_ARRAY_BUFFER, m_RendererID);
  }

  void unbind() {
    glBindBuffer(GL_ARRAY_BUFFER, 0);
  }

}
