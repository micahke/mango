package com.micahelias.scene;

import java.util.ArrayList;

public class SceneManager {
  
  // Singleton instance
  private static SceneManager instance;

  private Scene activeScene;
  private ArrayList<Scene> scenes;
  
  private SceneManager() {
    scenes = new ArrayList<Scene>();
  }

  public static SceneManager init() {
    if (instance == null) {
      instance = new SceneManager();
    }
    // TODO: Replace this with some sort of error checking or logging:
    return instance;
  }

  public Scene getActiveScene() {
    return instance.activeScene();
  }


  public Scene activeScene() {
    return this.activeScene;
  }

  public void setScene(Scene scene) {
    this.activeScene = scene;
  }

  public static SceneManager get() {
    return instance;
  }

  public ArrayList<Scene> listScenes() {
    return scenes;
  }


}
