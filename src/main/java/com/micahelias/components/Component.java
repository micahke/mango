package com.micahelias.components;

import com.micahelias.scene.Entity;

public class Component implements MScript {

  public Entity entity;
  public boolean enabled = false;

  public void enable() {
    this.enabled = true;
  }

  public void disable() {
    this.enabled = false;
  }

  // MSCRIPT STUFF
  public void init() {
  };

  public void update() {
  };

}
