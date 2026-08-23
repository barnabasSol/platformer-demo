package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

const (
	width      int32   = 800
	height     int32   = 450
	jump_speed float32 = -400
	hor_speed  float32 = 130
)

var floor int32

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(width, height, "platformer")

	gameMap, err := tiled.LoadFile("world.tmx")
	tilesetTexture := rl.LoadTexture("res/day_platformer/PNG/tileset.png")
	if err != nil {
		log.Fatal(err)
	}

	player := Player{
		FootPosition: rl.NewVector2(float32(width)/2, float32(height)/6),
		Grounded:     false,
		Gravity:      200,
		DragSpeed:    -340,
		FacingRight:  true,
		SpriteAnimation: &SpriteAnimation{
			StateTextures: map[string][]rl.Texture2D{},
			CurrentFrame:  0,
			FrameDuration: 0.099,
			Timer:         0,
			Rect:          rl.NewRectangle(0, 0, 512/2.4, 512/2.4),
			CurrentState:  StateJump,
			FramePinned:   true,
		},
	}
	player.SpriteAnimation.Load()
	player.JumpSpeed = player.Gravity

	cam := NewCamera(player)
	defer func() {
		rl.CloseWindow()
		player.SpriteAnimation.Unload()
	}()
	rl.SetTargetFPS(60)

	layer := gameMap.Layers[0]
	clouds := InitClouds(gameMap)
	trees := InitTrees()

	floor = trees.Texture.Height + 93
	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()

		player.Update(delta, []EnvironmentItem{})
		player.SpriteAnimation.Update(delta, &player)

		spriteCurrentState := player.SpriteAnimation.CurrentState
		currentFrame := player.SpriteAnimation.CurrentFrame
		currentTex := player.SpriteAnimation.StateTextures[spriteCurrentState]

		cam.Update(player, delta)

		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode2D(cam.Cam)
		for _, ct := range clouds {
			rl.DrawTexture(
				ct.Texture,
				int32(ct.X),
				int32(ct.Y),
				rl.White,
			)
		}
		trees.DrawTrees(gameMap)
		for y := 0; y < gameMap.Height; y++ {
			for x := 0; x < gameMap.Width; x++ {

				tile := layer.Tiles[y*gameMap.Width+x]

				if tile == nil {
					continue
				}

				localID := int(tile.ID)

				if localID == 0 {
					continue
				}

				tileX := localID % tilesetColumns
				tileY := localID / tilesetColumns

				source := rl.Rectangle{
					X:      float32(tileX * tileWidth),
					Y:      float32(tileY * tileHeight),
					Width:  tileWidth,
					Height: tileHeight,
				}

				destination := rl.Rectangle{
					X:      float32(x * tileWidth),
					Y:      float32(y * tileHeight),
					Width:  tileWidth,
					Height: tileHeight,
				}

				rl.DrawTexturePro(
					tilesetTexture,
					source,
					destination,
					rl.Vector2{},
					0,
					rl.White,
				)
			}
		}

		rl.DrawTexturePro(
			currentTex[currentFrame],
			player.SpriteAnimation.Src,
			player.SpriteAnimation.Dst,
			rl.Vector2{},
			0,
			rl.White,
		)
		rl.DrawCircleV(player.FootPosition, 3, rl.Red)
		rl.EndMode2D()
		rl.DrawFPS(20, 20)

		rl.EndDrawing()
	}
}
