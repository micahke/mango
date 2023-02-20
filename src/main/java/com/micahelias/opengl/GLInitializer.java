package com.micahelias.opengl;

import static org.lwjgl.glfw.GLFW.*;
import static org.lwjgl.opengl.GL30.*;


public final class GLInitializer {

  private static RenderAPI renderAPI;

  public static void setRenderAPI(RenderAPI api) {
    renderAPI = api;
    if (api == RenderAPI.OPENGL) {
      initGLFW(3, 3);
    }
  }


  private static void initGLFW(int major, int minor) {

    if (!glfwInit()) {
      System.out.println("Could not start GLFW");
      System.exit(-1);
    }

    glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, major);
    glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, minor);
    glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE);
    glfwWindowHint(GLFW_OPENGL_FORWARD_COMPAT, GL_TRUE);

  }


  public static RenderAPI getRenderAPI() {
    return renderAPI;
  }


}
