package com.micahelias.core;

import com.micahelias.opengl.GLInitializer;
import com.micahelias.opengl.RenderAPI;
import com.micahelias.scene.SceneManager;

import static org.lwjgl.glfw.GLFW.*;
import static org.lwjgl.opengl.GL30.*;

public final class Mango {

  public static Window window;
  public static Timer timer;
  public static SceneManager sceneManager;


  // Statically initialize the Mango engine
  public static void init() {
    GLInitializer.setRenderAPI(RenderAPI.OPENGL); 
    timer = Timer.init();
    sceneManager = SceneManager.init();
  } 


  public static void loop() {
    while (!window.shouldClose()) {

      timer.updateDeltaTime();

      glfwPollEvents();
      glClearColor(0.5f, 0.5f, 0.5f, 1.0f);
      glClear(GL_COLOR_BUFFER_BIT);

      sceneManager.getActiveScene().update();

      sceneManager.getActiveScene().render();

      glfwSwapBuffers(window.id());
    }
    glfwTerminate();
  }

  public static void createWindow(int width, int height, String title, boolean vSyncEnabled) {
    window = new Window(width, height, title, vSyncEnabled);
  }


}
