package com.micahelias.opengl;

import static org.lwjgl.opengl.GL11.*;
import static org.lwjgl.opengl.GL20.*;

import java.io.*;
import java.util.HashMap;

public class Shader {

  String m_VertexPath;
  String m_FragmentPath;
  int m_RendererID;
  HashMap<String, Integer> m_UniformLocationCache = new HashMap<>();


  public Shader(String vertexPath, String fragmentPath) {
    m_VertexPath = vertexPath;
    m_FragmentPath = fragmentPath;

    String vertexShader = parseShader(vertexPath);
    String fragmentShader = parseShader(fragmentPath);

    m_RendererID = createShader(vertexShader, fragmentShader);
  }


  public void bind() {
    glUseProgram(m_RendererID);
  }

  public void unbind() {
    glUseProgram(0);
  }




  public String parseShader(String path) {
    StringBuilder source = new StringBuilder();
    try {
      ClassLoader loader = Thread.currentThread().getContextClassLoader();
      InputStream is = loader.getResourceAsStream(path);
      BufferedReader reader = new BufferedReader(new InputStreamReader(is));
      String line;
      while ((line = reader.readLine()) != null) {
        source.append(line).append("//\n");
      }
      reader.close();
      return source.toString();
    } catch (IOException e) {
      System.err.println("Could not read file!");
      e.printStackTrace();
      System.exit(-1);
    }
    return "";
  }

  public int compileShader(String source, int type) {
    int id = glCreateShader(type);
    glShaderSource(id, source);
    glCompileShader(id);

    // TODO: Error handling

    int result = glGetShaderi(id, GL_COMPILE_STATUS);
    if (result == GL_FALSE) {
      String message = glGetShaderInfoLog(id);
      System.out.println("Failed to compile " + (type == GL_VERTEX_SHADER ? "vertex" : "fragment") + " shader");
      System.out.println(message);
      glDeleteShader(id);
      return 0;
    }

    return id;
  }

  public int createShader(String vertexShader, String fragmentShader) {
    int program = glCreateProgram();
    int vs = compileShader(vertexShader, GL_VERTEX_SHADER);
    int fs = compileShader(fragmentShader, GL_FRAGMENT_SHADER);

    glAttachShader(program, vs);
    glAttachShader(program, fs);
    glLinkProgram(program);
    glValidateProgram(program);

    glDeleteShader(vs);
    glDeleteShader(fs);


    return program;

  }


  // Set uniforms
  public void setUniform4f(String name, float v0, float v1, float v2, float v3) {
    glUniform4f(getUniformLocation(name), v0, v1, v2, v3);
  }

  public int getUniformLocation(String name) {
    if (m_UniformLocationCache.containsKey(name)) {
      return m_UniformLocationCache.get(name);
    }
    int location = glGetUniformLocation(m_RendererID, name);
    if (location == -1) {
      // assert("The uniform is not in the shader");
      System.out.println("Warning");
    }

    m_UniformLocationCache.put(name, location);
    return location;
  }


}
