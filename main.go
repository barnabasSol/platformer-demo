package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width      int32   = 800
	height     int32   = 450
	jump_speed float32 = -400
	hor_speed  float32 = 130
	floor      int32   = 340
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(width, height, "platformer")
	envItems := []EnvironmentItem{}
	game_map := Map{
		MaxWidth:  1600,
		MinWidth:  0,
		MaxHeight: 0,
		MinHeight: 1600,
	}
	game_map.initMap()

	player := Player{
		FootPosition: rl.NewVector2(float32(width)/2, float32(height)/2),
		Grounded:     false,
		Gravity:      200,
		DragSpeed:    -340,
		FacingRight:  true,
		SpriteAnimation: &SpriteAnimation{
			StateTextures: map[string][]rl.Texture2D{},
			CurrentFrame:  0,
			FrameDuration: 0.099,
			Timer:         0,
			Rect:          rl.NewRectangle(0, 0, 512/2, 512/2),
			CurrentState:  StateJump,
			FramePinned:   true,
		},
	}
	player.SpriteAnimation.Load()
	player.JumpSpeed = player.Gravity

	cam := NewCamera(player)
	_ = cam

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

		game_map.updateMap(player, delta)
		cam.Update(player.FootPosition, delta)

		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode2D(cam.Cam)
		for _, ei := range envItems {
			rl.DrawRectangleRec(ei.Rect, ei.Color)
		}
		game_map.drawMap()
		rl.DrawTexturePro(
			currentTex[currentFrame],
			player.SpriteAnimation.Src,
			player.SpriteAnimation.Dst,
			rl.Vector2{},
			0,
			rl.White,
		)
		rl.DrawCircleV(player.FootPosition, 3, rl.Pink)
		rl.DrawFPS(20, 20)
		rl.EndMode2D()

		rl.EndDrawing()
	}
}
