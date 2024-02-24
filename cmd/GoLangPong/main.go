package main

import (
	"fmt"

	"github.com/Ollie-Ave/GoLangPong/internal/constants"
	"github.com/Ollie-Ave/GoLangPong/internal/entities"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
    fmt.Println("Hello, World!")

    rl.InitWindow(constants.WindowWidth, constants.WindowHeight, constants.WindowTitle)
    defer rl.CloseWindow()

    rl.SetExitKey(constants.WindowExitKey)
    rl.SetTargetFPS(constants.WindowTargetFPS)

    entityManager := entities.NewEntityManager()

    paddle1 := entities.NewPaddle(rl.NewVector2(20, 20), entities.NewPaddleControls(rl.KeyW, rl.KeyS))
    entityManager.SpawnEntity(paddle1)

    paddle2 := entities.NewPaddle(rl.NewVector2(float32(rl.GetScreenWidth() - (20 + constants.PaddleWidth)), 20), entities.NewPaddleControls(rl.KeyUp, rl.KeyDown))
    entityManager.SpawnEntity(paddle2)

    for !rl.WindowShouldClose() {
        update(entityManager)
    }
}

func update(entityManager entities.IEntityManager) {
    rl.BeginDrawing()
    rl.ClearBackground(constants.WindowBackgroundColor)
    
    for _, entity := range entityManager.GetEntities() {
        entity.Update()

        if  renderableEntity, isRenderable := entity.(entities.IRenderable); isRenderable {
            renderableEntity.Render()
        }
    }
    
    rl.EndDrawing()
}
