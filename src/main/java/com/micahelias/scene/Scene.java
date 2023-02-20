package com.micahelias.scene;

import java.util.ArrayList;

import com.micahelias.components.MeshRenderer;

public class Scene {

  public String name;
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
