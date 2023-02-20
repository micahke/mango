package com.micahelias;

import com.micahelias.components.MeshRenderer;
import com.micahelias.core.Mango;
import com.micahelias.graphic.Triangle2D;
import com.micahelias.scene.Entity;
import com.micahelias.scene.Scene;
import com.micahelias.util.Color;

public class App {


  public static void main(String[] args) {
    Mango.init();
    Mango.createWindow(800, 600, "Mango", true);
    Scene mainScene = new Scene("home");
    mainScene.setBackgroundColor(Color.DRACULA);
    Mango.sceneManager.setScene(mainScene);

    Entity character = new Entity("mainCharacter");
    character.addComponent(new Triangle2D().setColor(Color.MINT_LEAF));
    character.addComponent(new MeshRenderer());
    mainScene.addEntity(character);

    Mango.loop();
  }


}
