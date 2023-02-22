package com.micahelias.scene;

import org.joml.Matrix4f;

import com.micahelias.core.Mango;

public class Camera2D extends Entity {

  public Matrix4f projectionMatrix;
  public Matrix4f viewMatrix;

  public Camera2D() {
    super("camera2D");

    int windowWidth = Mango.window.width();
    int windowHeight = Mango.window.height();

    this.projectionMatrix = new Matrix4f().ortho(0, windowWidth, 0, windowHeight, -1.0f, 1.0f);
    this.viewMatrix = new Matrix4f().translate(0.0f, 0.0f, 0.0f);

  }

  public Matrix4f projectionMatrix() {
    return this.projectionMatrix;
  }

  public Matrix4f viewMatrix() {
    return this.viewMatrix;
  }

}
