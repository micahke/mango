package com.micahelias.core;

import static org.lwjgl.glfw.GLFW.*;

public class Timer {

  // Singleton instance
  private static Timer instance;

  private float deltaTime;
  private float lastFrame = 0;

  private double timeDilation = 1;

  private Timer() {}

  public static Timer init() {
    if (instance == null) {
      instance = new Timer();
    }
    return instance;
  }


  public static Timer get() {
    return instance;
  }


  public float getTime() {
    return (float)glfwGetTime();
  }

  public void setTimeDilation(double timeDilation) {
    this.timeDilation = timeDilation;
  }

  void updateDeltaTime() {
    deltaTime = getTime() - lastFrame;
    lastFrame = getTime();
  }

  public float deltaTime() {
    return deltaTime * (float)timeDilation;
  }

}

