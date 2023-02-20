package com.micahelias.opengl;

import java.nio.ByteBuffer;

public class Texture {


  int m_RendererID;
  String m_FilePath;
  ByteBuffer m_LocalBuffer;
  int m_Width, m_Height, m_BPP;


  public Texture() {

  }
  
  void bind() {

  }

  void unbind() {

  }
  
  int getWidth() { return m_Width; }
  int getHeight() { return m_Height; }


}
