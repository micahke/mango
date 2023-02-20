package com.micahelias.opengl;



import static org.lwjgl.opengl.GL20.*;
import static org.lwjgl.opengl.GL30.*;

import java.util.ArrayList;


public class VertexArray {

  int m_RendererID;

  public VertexArray() {
    m_RendererID = glGenVertexArrays();
  }

  public void addBuffer(VertexBuffer vb, VertexBufferLayout layout) {

    bind();
    vb.bind();

    ArrayList<VertexBufferElement> elements = layout.getElements();
    int offset = 0;
    for (int i = 0; i < elements.size(); i++) {
      VertexBufferElement element = elements.get(i);
      glEnableVertexAttribArray(i);
      glVertexAttribPointer(i, element.count, element.type, element.normalized, layout.getStride(), offset);
      offset += element.getTypeSize();
    }

  }

  public void bind() {
    glBindVertexArray(m_RendererID);
  }

  public void unbind() {
    glBindVertexArray(0);
  }


}

