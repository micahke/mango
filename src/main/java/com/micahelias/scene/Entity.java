package com.micahelias.scene;

import java.util.ArrayList;

import com.micahelias.components.Component;
import com.micahelias.components.TransformComponent;
import com.micahelias.core.Mango;

public class Entity {

  public String name;
  public Scene scene;
  public TransformComponent transform;
  
  private ArrayList<Component> components;

  public Entity(String name) {
    this.name = name;
    this.components = new ArrayList<Component>();
    this.addComponent(new TransformComponent(0, 0));
  }

  public void setScene(Scene scene) {
    this.scene = scene;
  }

  public String name() {
    return this.name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public ArrayList<Component> getComponents() {
    return this.components;
  }

  
  public void addComponent(Component c) {
    c.entity = this;
    components.add(c);
    if (TransformComponent.class.isInstance(c)) {
      transform = TransformComponent.class.cast(c);
    }
  }

  public void update() {
    for (Component c : components) {
      if (!c.enabled) {
        c.init();
        c.enable();
      }
      c.update();
    }
  }

  public <T extends Component> T getComponent(Class<T> componentClass) {
    for (Component c : this.components) {
      if (componentClass.isInstance(c)) {
        return componentClass.cast(c);
      }
    }
    return null;
  }

  // Can be used from scripts to find an entity within the game
  public static Entity find(String name) {
    for (Entity entity : Mango.sceneManager.getActiveScene().getEntities()) {
      if (entity.name.equals(name)) {
        return entity;
      }
    }
    return null;
  }


}
