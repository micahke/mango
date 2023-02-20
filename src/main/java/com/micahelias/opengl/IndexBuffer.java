package com.micahelias.opengl;


import static org.lwjgl.opengl.GL15.*;

/**
 * VertexBuffer
 */
public class IndexBuffer {

  private int m_RendererID;
  private int m_Count;
  
  public IndexBuffer(int[] data) {
    m_Count = data.length;
    m_RendererID = glGenBuffers();
    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, m_RendererID);
    glBufferData(GL_ELEMENT_ARRAY_BUFFER, data, GL_STATIC_DRAW);
  }


  public void bind() {
    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, m_RendererID);
  }

  public void unbind() {
    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, 0);
  }

  public int getCount() {
    return m_Count;
  }

}



