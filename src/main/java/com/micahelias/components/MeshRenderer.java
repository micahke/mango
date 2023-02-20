package com.micahelias.components;

import static org.lwjgl.opengl.GL30.*;


public class MeshRenderer extends Component {

  ModelComponent model;

  public void render() {
    model = entity.getComponent(ModelComponent.class);
    if (model == null) {
      return;
    };

    if (model.ibo != null) {
      model.ibo.bind();
    }
    model.vao.bind();
    model.shader.bind();

    if (model.ibo != null) {
      glDrawElements(GL_TRIANGLES, model.ibo.getCount(), GL_UNSIGNED_INT, 0);
    }

  }

  public void init() {};
  public void update() {};


}

  
