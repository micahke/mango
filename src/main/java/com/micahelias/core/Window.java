package com.micahelias.core;

import static org.lwjgl.glfw.GLFW.*;
import org.lwjgl.opengl.GL;

public class Window {

  private long id;

  private int height;
  private int width;
  private String title;
  private boolean vSyncEnabled;

  // TODO: make this public when it moves out of the core package
  Window(int width, int height, String title, boolean vSyncEnabled) {
    long windowID = glfwCreateWindow(width, height, title, 0, 0);
    if (windowID == 0) {
      glfwTerminate();
      System.out.println("Could not create GLFW window");
      System.exit(1);
    }
    glfwMakeContextCurrent(windowID);
    if (vSyncEnabled) enableVsync();
    else disableVsync();
    GL.createCapabilities(); // crucial

    this.id = windowID;
    this.width = width;
    this.height = height;
    this.title = title;
    this.vSyncEnabled = vSyncEnabled;

  }

  public long id() {
    return id;
  }

  public void enableVsync(){
    glfwSwapInterval(1);
    this.vSyncEnabled = true;
  }

  public void disableVsync(){
    glfwSwapInterval(0);
    this.vSyncEnabled = false;
  }

  public int width() {
    return this.width;
  }

  public int height() {
    return this.height;
  }

  public String title() {
    return this.title;
  }

  public boolean vSyncEnabled() {
    return this.vSyncEnabled;
  }

  public boolean shouldClose() {
    return glfwWindowShouldClose(this.id);
  }


}
