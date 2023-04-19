package main

import (
	"math"
	"math/rand"
	"runtime"

	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/core"
	"github.com/micahke/infinite-universe/mango/im"
	"github.com/micahke/infinite-universe/mango/logging"
	"github.com/micahke/infinite-universe/mango/util/color"
)

func init() {
  runtime.LockOSThread()
}

const (
  WIDTH int = 800
  HEIGHT int = 600

  PADDLE_WIDTH = 25
  PADDLE_HEIGHT = 100

  BALL_SIZE = 20
  MAX_Y_VELO = 175

  PADDLE_SPEED = 250
)


func main() {

  mango.Init(core.RENDER_MODE_IM)
  mango.CreateWindow(WIDTH, HEIGHT, "Pong", true)

  mango.IM.ConnectScene(&Pong{})
  mango.IM.SetBackgroundColor(color.DRACULA)
  mango.Start()

}


type Pong struct {

  leftPaddle, rightPaddle Paddle
  ball Ball

}

type Paddle struct {
  x, y float32
  width, height float32
  targetHitPoint float32
}

type Ball struct {
  x, y float32
  xVelo, yVelo float32
}

// Mango required functions for IM

func (pong *Pong) Init() {
  pong.leftPaddle = Paddle{
    x: 0,
    y: float32(HEIGHT) / 2 - float32(PADDLE_HEIGHT / 2),
    width: PADDLE_WIDTH,
    height: PADDLE_HEIGHT,
  }

  pong.rightPaddle = Paddle{
    x: float32(WIDTH) - PADDLE_WIDTH,
    y: float32(HEIGHT) / 2 - float32(PADDLE_HEIGHT / 2),
    width: PADDLE_WIDTH,
    height: PADDLE_HEIGHT,
  }

  pong.ball = Ball{
    x: float32(WIDTH) / 2 - (float32(BALL_SIZE) / 2),
    y: float32(HEIGHT) / 2 - (float32(BALL_SIZE) / 2),
    xVelo: -300,
  }

  pong.leftPaddle.calcTarget()
}

func (pong *Pong) Update(deltaTime float64) {
  pong.updateBall();

    pong.leftPaddle.upOrDown(pong.ball)
    pong.rightPaddle.upOrDown(pong.ball)

}

func (pong *Pong) Draw() {
  pong.drawPaddles();
  pong.drawBall();
}

// +++++++++++++++++++++++++++++++++++++++++++++++


func (pong *Pong) drawPaddles() {

  // Using batched quads
  q1 := im.Quad{
    X: pong.leftPaddle.x,
    Y: pong.leftPaddle.y,
    Width: PADDLE_WIDTH,
    Height: PADDLE_HEIGHT,
    Color: color.MINT_GREEN,
  }
  q2 := im.Quad{
    X: pong.rightPaddle.x,
    Y: pong.rightPaddle.y,
    Width: PADDLE_WIDTH,
    Height: PADDLE_HEIGHT,
    Color: color.MINT_GREEN,
  }

  mango.IM.DrawQuad(q1)
  mango.IM.DrawQuad(q2)

}


func (pong *Pong) drawBall() {

  mango.IM.DrawCircle(pong.ball.x, pong.ball.y, BALL_SIZE, BALL_SIZE, color.WHITE)

}


func (pong *Pong) updateBall() {

  pong.handleWallCollision()
  // only check collisions when ball is getting close to edge
  if pong.ball.x <= PADDLE_WIDTH && pong.ball.xVelo < 0  {
    pong.handleLeftCollision()
  }

  if pong.ball.x + BALL_SIZE >= float32(WIDTH) - PADDLE_WIDTH && pong.ball.xVelo > 0 {
    pong.handleRightCollision()
  }

  // check for wall collisions

  pong.ball.x += pong.ball.xVelo * float32(core.Timer.DeltaTime())
  pong.ball.y += pong.ball.yVelo * float32(core.Timer.DeltaTime())


}

func (pong *Pong) handleLeftCollision() {
  if pong.ball.y <= pong.leftPaddle.y + PADDLE_HEIGHT && pong.ball.y + BALL_SIZE >= pong.leftPaddle.y {
    pong.ball.xVelo = -pong.ball.xVelo
    pong.calculateYVelo(pong.leftPaddle)
    pong.rightPaddle.calcTarget()
  }
}

func (pong *Pong) handleRightCollision() {
  if pong.ball.y <= pong.rightPaddle.y + PADDLE_HEIGHT && pong.ball.y + BALL_SIZE >= pong.rightPaddle.y {
    pong.ball.xVelo = -pong.ball.xVelo
    pong.calculateYVelo(pong.rightPaddle)
    pong.leftPaddle.calcTarget()
  }
}


func (pong *Pong) calculateYVelo(paddle Paddle) {

  paddleCenter := paddle.y + (PADDLE_HEIGHT / 2)
  ballCenter := pong.ball.y + (BALL_SIZE / 2)

  distance := paddleCenter - ballCenter

  percentageOfMax := float32(distance) / float32(PADDLE_HEIGHT / 2)

  pong.ball.yVelo = -percentageOfMax * MAX_Y_VELO

}


func (pong *Pong) handleWallCollision() {

  if pong.ball.y + BALL_SIZE >= float32(HEIGHT){
    pong.ball.yVelo = float32(-1 * math.Abs(float64(pong.ball.yVelo)));
  }
if pong.ball.y  <= 0 {
    pong.ball.yVelo = float32(math.Abs(float64(pong.ball.yVelo)));
  }
}



// ++++++++++++++++++++++++++++++++++++++++++++++++++++++


// calculates the next target hit point relative to center
func (paddle *Paddle) calcTarget() {

  // pick whether up or down
  dir := rand.Intn(1)

  targetHitPointRaw := rand.Intn(PADDLE_HEIGHT / 2)

  if dir == 0 {
    targetHitPointRaw = -1 * targetHitPointRaw
  }

  paddle.targetHitPoint = float32(targetHitPointRaw)

}

func (paddle *Paddle) upOrDown(ball Ball) {

  paddleCenter := paddle.y + (PADDLE_HEIGHT / 2)

  targetPoint := paddleCenter + paddle.targetHitPoint

  logging.Log(paddle.targetHitPoint)


  if targetPoint > ball.y + BALL_SIZE {

    paddle.y -= float32(PADDLE_SPEED * core.Timer.DeltaTime())

  }
  if targetPoint < ball.y + BALL_SIZE {

    paddle.y += float32(PADDLE_SPEED * core.Timer.DeltaTime())


  }

}
