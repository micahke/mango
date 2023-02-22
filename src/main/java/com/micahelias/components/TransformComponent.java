package com.micahelias.components;

import org.joml.Vector2f;
import org.joml.Vector3f;

public class TransformComponent extends Component {

  public Vector3f position;

  public TransformComponent(float x, float y, float z) {
    position = new Vector3f(x, y, z);
  }

  public TransformComponent(float x, float y) {
    position = new Vector3f(x, y, 0.0f);
  }

  public TransformComponent(Vector3f position) {
    this.position = position;
  }

  public TransformComponent(Vector2f position) {
    this.position = new Vector3f(position.x, position.y, 0.0f);
  }

  public void init() {}

  public void update() {
  }


}
