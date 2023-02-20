package com.micahelias.scene;

import static org.lwjgl.opengl.GL30.*;
import java.util.ArrayList;

import org.joml.Vector4f;

import com.micahelias.components.MeshRenderer;
import com.micahelias.util.Color;

public class Scene {

  public String name;
  public Color backgroundColor;
  private ArrayList<Entity> entities;
  

  public Scene(String name) {
    this.name = name;
    this.entities = new ArrayList<Entity>();
  }

  public void addEntity(Entity e) {
    e.setScene(this);
    entities.add(e);
  }

  public void update() {
    for (Entity e : entities) {
      e.update();
    }
  }

  public void setBackgroundColor(Color color) {
    this.backgroundColor = color;
  }

  public void clear() {
    Vector4f colorVec = backgroundColor.toVector();
    glClearColor(colorVec.x, colorVec.y, colorVec.z, colorVec.w);
    glClear(GL_COLOR_BUFFER_BIT);
  }

  public void render() {
    for (Entity e : entities) {
      MeshRenderer renderer = e.getComponent(MeshRenderer.class);
      if (renderer != null) {
        renderer.render();
      }
    }
  }

  public String getName() {
    return this.name;
  }

  public ArrayList<Entity> getEntities() {
    return this.entities;
  }

}
