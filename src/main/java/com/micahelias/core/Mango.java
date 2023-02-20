package com.micahelias.core;

import com.micahelias.opengl.GLInitializer;
import com.micahelias.opengl.RenderAPI;
import static org.lwjgl.glfw.GLFW.*;
import static org.lwjgl.opengl.GL30.*;

public final class Mango {

  public static Window window;
  public static Timer timer;


  // Statically initialize the Mango engine
  public static void init() {
    GLInitializer.setRenderAPI(RenderAPI.OPENGL); 
    timer = new Timer();
  } 


  public static void loop() {
    while (!window.shouldClose()) {

      timer.updateDeltaTime();

      glfwPollEvents();
      glClearColor(0.5f, 0.5f, 0.5f, 1.0f);
      glClear(GL_COLOR_BUFFER_BIT);


      glfwSwapBuffers(window.id());
    }
    glfwTerminate();
  }

  public static void createWindow(int width, int height, String title, boolean vSyncEnabled) {
    window = new Window(width, height, title, vSyncEnabled);
  }


}
