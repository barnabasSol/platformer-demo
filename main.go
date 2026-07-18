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

var gravity float32 = 200
var drag_speed float32 = -340

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(width, height, "platformer")
	defer rl.CloseWindow()

	player := Player{
		Position:    rl.NewVector2(float32(width)/2, float32(height)/2),
		CanJump:     false,
		FacingRight: true,
		SpriteAnimation: &SpriteAnimation{
			StateTextures: map[string][]rl.Texture2D{},
			CurrentFrame:  0,
			FrameDuration: 0.099,
			Timer:         0,
			CurrentState:  "idle",
		},
	}
	player.SpriteAnimation.Load()
	player.Rect.Width = 512 / 2
	player.Rect.Height = 512 / 2
	player.Speed = gravity

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()
		player.Update(delta)
		player.SpriteAnimation.Update(delta, &player)
		spriteCurrentState := player.SpriteAnimation.CurrentState
		currentFrame := player.SpriteAnimation.CurrentFrame
		currentTex := player.SpriteAnimation.StateTextures[spriteCurrentState]

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		// rl.DrawRectangleRec(player.Rect, rl.White)

		rl.DrawTexturePro(
			currentTex[currentFrame],
			player.SpriteAnimation.Src,
			player.SpriteAnimation.Dst,
			rl.Vector2{},
			0,
			rl.White,
		)
		rl.DrawCircleV(player.Position, 3, rl.Pink)

		rl.EndDrawing()
	}
}
