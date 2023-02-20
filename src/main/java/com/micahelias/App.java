package com.micahelias;

import com.micahelias.core.Mango;

public class App {


  public static void main(String[] args) {
    Mango.init();
    Mango.createWindow(800, 600, "Mango", true);
    Mango.timer.deltaTime();
    Mango.loop();
  }


}
