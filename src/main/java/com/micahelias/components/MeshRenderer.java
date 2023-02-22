package com.micahelias.components;

import static org.lwjgl.opengl.GL30.*;

import org.joml.Matrix4f;
import org.joml.Vector3f;

import com.micahelias.scene.Camera2D;
import com.micahelias.scene.Entity;


public class MeshRenderer extends Component {

  ModelComponent model;

  public void render() {
    model = entity.getComponent(ModelComponent.class);
    Camera2D camera = (Camera2D)Entity.find("camera2D"); // find the active camera in the scene
    
    if (camera == null) {
      System.out.println("No camera found scene");
    }

    if (model == null) {
      return;
    };

    if (model.ibo != null) {
      model.ibo.bind();
    }
    model.vao.bind();
    model.shader.bind();

    // Set model color
    model.shader.setUniform4f("u_Color", model.color.toVector());
    model.shader.setUniformMat4f("projectionMatrix", camera.projectionMatrix());

    Vector3f position = entity.transform.position;
    Matrix4f translation = new Matrix4f().translate(position.x, position.y, 0); 
    translation = translation.mul(model.getModelMatrix());
    model.shader.setUniformMat4f("modelMatrix", translation);

    if (model.ibo != null) {
      glDrawElements(GL_TRIANGLES, model.ibo.getCount(), GL_UNSIGNED_INT, 0);
    }

  }

  public void init() {};
  public void update() {};


}

  
