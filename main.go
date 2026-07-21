package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width      int32   = 800
	height     int32   = 450
	jump_speed float32 = -400
	hor_speed  float32 = 130
	floor      int32   = height
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(width, height, "platformer")
	envItems := []EnvironmentItem{
		{Rect: rl.NewRectangle(50, 380, 150, 10), Color: rl.Black},
		{Rect: rl.NewRectangle(350, 320, 150, 10), Color: rl.Black},
	}

	player := Player{
		Position:    rl.NewVector2(float32(width)/2, float32(height)/2),
		Grounded:    false,
		Gravity:     200,
		DragSpeed:   -340,
		FacingRight: true,
		SpriteAnimation: &SpriteAnimation{
			StateTextures: map[string][]rl.Texture2D{},
			CurrentFrame:  0,
			FrameDuration: 0.099,
			Timer:         0,
			Rect:          rl.NewRectangle(0, 0, 512/2, 512/2),
			CurrentState:  "jump",
			FramePinned:   true,
		},
	}
	player.SpriteAnimation.Load()
	player.JumpSpeed = player.Gravity

	defer func() {
		rl.CloseWindow()
		player.SpriteAnimation.Unload()
	}()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()
		player.Update(delta, envItems)
		player.SpriteAnimation.Update(delta, &player)
		spriteCurrentState := player.SpriteAnimation.CurrentState
		currentFrame := player.SpriteAnimation.CurrentFrame
		currentTex := player.SpriteAnimation.StateTextures[spriteCurrentState]

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)

		for _, ei := range envItems {
			rl.DrawRectangleRec(ei.Rect, ei.Color)
		}
		rl.DrawTexturePro(
			currentTex[currentFrame],
			player.SpriteAnimation.Src,
			player.SpriteAnimation.Dst,
			rl.Vector2{},
			0,
			rl.White,
		)
		rl.DrawCircleV(player.Position, 3, rl.Pink)
		rl.DrawFPS(20, 20)

		rl.EndDrawing()
	}
}
