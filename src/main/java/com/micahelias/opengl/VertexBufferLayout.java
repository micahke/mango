package com.micahelias.opengl;

import java.util.ArrayList;
import static org.lwjgl.opengl.GL11.*;

class VertexBufferElement {
  int type;
  int count;
  boolean normalized;

  public VertexBufferElement(int type, int count, boolean normalized) {
    this.type = type;
    this.count = count;
    this.normalized = normalized;
  }

  public int getTypeSize() {
    switch(type){
      case GL_FLOAT:
        return Float.BYTES;
      case GL_UNSIGNED_INT:
        return Integer.BYTES;
      case GL_UNSIGNED_BYTE:
        return Character.BYTES;
      default:
      return 0;
    }
  }

}

public class VertexBufferLayout {

  ArrayList<VertexBufferElement> m_Elements = new ArrayList<VertexBufferElement>();
  int m_Stride = 0;

  public void pushf(int count) {
    m_Elements.add(new VertexBufferElement(GL_FLOAT, count, false));
    m_Stride += Float.BYTES * count;
  }

  public void pushi(int count) {
    m_Elements.add(new VertexBufferElement(GL_UNSIGNED_INT, count, false));
    m_Stride += Integer.BYTES * count;
  }


  public void pushc(int count) {
    m_Elements.add(new VertexBufferElement(GL_UNSIGNED_BYTE, count, true));
    m_Stride += Character.BYTES * count;
  }

  public ArrayList<VertexBufferElement> getElements() {
    return m_Elements;
  }

  public int getStride() {
    return m_Stride;
  }



}
