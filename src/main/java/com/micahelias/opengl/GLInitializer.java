package com.micahelias.opengl;

import static org.lwjgl.glfw.GLFW.*;
import static org.lwjgl.opengl.GL30.*;


public final class GLInitializer {

  private static RenderAPI renderAPI;

  public static void setRenderAPI(RenderAPI api) {
    renderAPI = api;
    if (api == RenderAPI.OPENGL) {
      initGLFW();
    }
  }


  private static void initGLFW() {

    if (!glfwInit()) {
      System.out.println("Could not start GLFW");
      System.exit(-1);
    }

    glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3);
    glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 2);
    glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE);
    glfwWindowHint(GLFW_OPENGL_FORWARD_COMPAT, GL_TRUE);

  }




}
