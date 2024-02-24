package entities

import (
	"github.com/Ollie-Ave/GoLangPong/internal/constants"
	"github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
    position rl.Vector2
    colour rl.Color

    controls *PaddleControls
}

type PaddleControls struct {
    upKey int32
    downKey int32
}

func NewPaddleControls(upKey int32, downKey int32) *PaddleControls {
    return &PaddleControls {
        upKey: upKey,
        downKey: downKey,
    }
}

func NewPaddle(initialPosition rl.Vector2, controls *PaddleControls) IEntity {
    return &Paddle {
        position: initialPosition,
        colour: rl.White,
        controls: controls,
    }
}

func (paddle *Paddle) Update() {
    paddle.handleInput()
}

func (paddle *Paddle) Render() {
    rl.DrawRectangleV(paddle.position, rl.NewVector2(constants.PaddleWidth, constants.PaddleHeight), paddle.colour)
}

func (paddle *Paddle) handleInput() {
    if rl.IsKeyDown(paddle.controls.upKey) && paddle.position.Y > 0 {
        paddle.position.Y -= constants.PaddleSpeed
    }

    if rl.IsKeyDown(paddle.controls.downKey) && paddle.position.Y < float32(rl.GetScreenHeight()) - constants.PaddleHeight {
        paddle.position.Y += constants.PaddleSpeed
    }
}
