package com.micahelias;

import com.micahelias.components.Component;
import com.micahelias.components.MeshRenderer;
import com.micahelias.core.Mango;
import com.micahelias.graphic.Rect2D;
import com.micahelias.scene.Camera2D;
import com.micahelias.scene.Entity;
import com.micahelias.scene.Scene;
import com.micahelias.util.Color;

public class App {

  static int WIDTH = 800;
  static int HEIGHT = 600;

  static int PADDLE_WIDTH = 30;
  static int PADDLE_HEIGHT = 100;

  static int BALL_SIZE = 20;

  public static void main(String[] args) {


    Mango.init();
    Mango.createWindow(WIDTH, HEIGHT, "GPU Pong", false);

    Scene scene = new Scene("main");
    scene.setBackgroundColor(Color.DRACULA);
    scene.addEntity(new Camera2D());
    Mango.sceneManager.setScene(scene);
    Mango.timer.setTimeDilation(10);

    Entity midline = new Entity("midline");
    midline.addComponent(new Rect2D(2, HEIGHT).setColor(new Color(69, 79, 82)));
    midline.addComponent(new MeshRenderer());
    midline.transform.position.x = 399.0f;



    Entity p1 = new Entity("player1");
    p1.addComponent(new Rect2D(PADDLE_WIDTH, PADDLE_HEIGHT).setColor(Color.MINT_LEAF));
    p1.addComponent(new MeshRenderer());
    p1.addComponent(new PaddleFollowComponent(20));
    p1.transform.position.y = (HEIGHT / 2) - (PADDLE_HEIGHT / 2);

    Entity p2 = new Entity("player2");
    p2.addComponent(new Rect2D(PADDLE_WIDTH, PADDLE_HEIGHT).setColor(Color.MINT_LEAF));
    p2.addComponent(new MeshRenderer());
    p2.addComponent(new PaddleFollowComponent(20));
    p2.transform.position.x = WIDTH - PADDLE_WIDTH;
    p2.transform.position.y = (HEIGHT / 2) - (PADDLE_HEIGHT / 2) + 40;

    Entity ball = new Entity("ball");
    ball.addComponent(new Rect2D(BALL_SIZE, BALL_SIZE).setColor(Color.PINK_GLAMOUR));
    ball.addComponent(new MeshRenderer());
    ball.addComponent(new BallController(BALL_SIZE));
    ball.transform.position.x = (WIDTH / 2) - (BALL_SIZE / 2);
    ball.transform.position.y = (HEIGHT / 2) - (BALL_SIZE / 2);

    scene.addEntity(midline);
    scene.addEntity(p1);
    scene.addEntity(p2);
    scene.addEntity(ball);
    Mango.loop();

  }
}




class BallController extends Component{

  int BALL_SIZE;
  Entity paddle1;
  Entity paddle2;

  int p1Score = 0;
  int p2Score = 0;

  int xVelo;
  float yVelo;


  public BallController(int BALL_SIZE) {
    this.BALL_SIZE = BALL_SIZE;
  }

  public void init() {
    paddle1 = Entity.find("player1");
    paddle2 = Entity.find("player2");


    xVelo = -1;
    yVelo = 0;
  }

  void resetGame() {
    System.out.println("Left  " + p1Score + " - " + p2Score + "  Right");
    float paddlesY = (600 / 2) - 50;
    paddle1.transform.position.y = paddlesY;
    paddle2.transform.position.y = paddlesY;
    entity.transform.position.x = (800 / 2) - 10;
    entity.transform.position.y = (600 / 2) - 10;
    xVelo = -1;
    yVelo = 0;
  }

  public void update() {
    checkPosition();
    entity.transform.position.x += xVelo * 30 * Mango.timer.deltaTime();
    entity.transform.position.y += yVelo * 20 * Mango.timer.deltaTime();
    if (entity.transform.position.x + BALL_SIZE <= 0) {
      p2Score++;
      resetGame();
    }
    if (entity.transform.position.x >= 800) {
      p1Score++;
      resetGame();
    }
  }

  public void checkPosition() {
    if (xVelo < 0 && entity.transform.position.x <= 30) {
      checkCollision(paddle1);
    }
    if (xVelo > 0 && entity.transform.position.x + 20 >= Mango.window.width() - 30) {
      checkCollision(paddle2);
    }
    if (entity.transform.position.y + BALL_SIZE >= 600 || entity.transform.position.y <= 0) {
      yVelo *= -1;
    }
  }

  public void checkCollision(Entity paddle) {
    float paddleY = paddle.transform.position.y;
    if (entity.transform.position.y + BALL_SIZE >= paddleY && entity.transform.position.y <= paddleY + 100) {
      xVelo *= -1;
      calcYVelo(paddle);
    }
  }



  public void calcYVelo(Entity e) {
    // 1) find the distance from the center
    float paddleCenter = e.transform.position.y + 50;
    float ballCenter = entity.transform.position.y + (BALL_SIZE / 2);
    float distance = -1 * (paddleCenter - ballCenter);

    yVelo = (float)distance / 50.0f;
  }

}


class PaddleFollowComponent extends Component {

  double speed;
  float hitPoint = 0;
  Entity ball;

  public PaddleFollowComponent(double speed) {
    super();
    this.speed = speed;
  }


  public void init() {
    ball = Entity.find("ball");
  }


  public void update() {
    if (ball.transform.position.x >= 395 && ball.transform.position.x <= 405) generateHitpoint();
    if (distanceFromBall() > 0) {
      entity.transform.position.y -= speed * Mango.timer.deltaTime();
    }
    if (distanceFromBall() < 0) {
      entity.transform.position.y += speed * Mango.timer.deltaTime();
    }
    if (entity.transform.position.y <= 0) entity.transform.position.y = 0;
    if (entity.transform.position.y + 100 >= 600) entity.transform.position.y = 500;
  }

  float distanceFromBall() {
    float ballCenter = ball.transform.position.y + 10;
    float paddleCenter = entity.transform.position.y + 50 + hitPoint;
    return paddleCenter - ballCenter;
  }

  void generateHitpoint() {
    hitPoint = (float)(Math.random() * 121) - 60;
  }


}
