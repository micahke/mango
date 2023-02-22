package com.micahelias;

import org.joml.Vector2f;

import com.micahelias.components.MeshRenderer;
import com.micahelias.core.Mango;
import com.micahelias.graphic.Rect2D;
import com.micahelias.graphic.Triangle2D;
import com.micahelias.scene.Camera2D;
import com.micahelias.scene.Entity;
import com.micahelias.scene.Scene;
import com.micahelias.util.Color;

public class App {


  public static void main(String[] args) {
    Mango.init();
    Mango.createWindow(800, 600, "Mango", true);
    Scene mainScene = new Scene("home");
    mainScene.setBackgroundColor(Color.DRACULA);
    mainScene.addEntity(new Camera2D());
    Mango.sceneManager.setScene(mainScene);

    Entity character = new Entity("mainCharacter");
    character.addComponent(new Triangle2D(
      new Vector2f(0.0f, 0),
      new Vector2f(100, 100),
      new Vector2f(250, 150)
    ).setColor(Color.MINT_LEAF));
    character.addComponent(new MeshRenderer());

    Entity rect = new Entity("rect");
    rect.transform.position.x = 500;
    rect.addComponent(new Rect2D(100, 100));
    rect.addComponent(new MeshRenderer());

    mainScene.addEntity(character);
    mainScene.addEntity(rect);

    Mango.loop();
  }


}
