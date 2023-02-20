package com.micahelias.core;

import static org.lwjgl.glfw.GLFW.*;

public class Timer {


  private float deltaTime;
  private float lastFrame = 0;


  public float getTime() {
    return (float)glfwGetTime();
  }


  void updateDeltaTime() {
    deltaTime = getTime() - lastFrame;
    lastFrame = getTime();
  }

  public float deltaTime() {
    return deltaTime;
  }

}

