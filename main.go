package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

const (
	width      int32   = 800
	height     int32   = 450
	jump_speed float32 = -600
	hor_speed  float32 = 130
)

func main() {

	rl.InitWindow(width, height, "platformer")

	gameMap, err := tiled.LoadFile("world.tmx")
	var envObjs []EnvironmentObject
	for _, objectLayer := range gameMap.ObjectGroups {
		if objectLayer.Name != "Collision" {
			continue
		}

		for _, object := range objectLayer.Objects {
			envObjs = append(envObjs, EnvironmentObject{
				Rect: rl.Rectangle{
					X:      float32(object.X),
					Y:      float32(object.Y),
					Width:  float32(object.Width),
					Height: float32(object.Height),
				},
				Color: rl.Red,
			})
		}
	}
	tilesetTexture := rl.LoadTexture("res/day_platformer/PNG/tileset.png")
	if err != nil {
		log.Fatal(err)
	}
	defer rl.UnloadTexture(tilesetTexture)

	playerPos := rl.NewVector2(float32(width)/1.8, float32(height)/6)
	var players []Player
	player := Player{
		Grounded:    false,
		Gravity:     200,
		DragSpeed:   -340,
		FacingRight: true,
		Box: rl.NewRectangle(
			playerPos.X-13,
			playerPos.Y-60,
			30,
			62,
		),
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
	players = append(players, player)
	antagonist := Player{
		Grounded:    false,
		Gravity:     200,
		DragSpeed:   -340,
		FacingRight: true,
		Box: rl.NewRectangle(
			playerPos.X-13,
			playerPos.Y-60,
			30,
			62,
		),
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
	antagonist.SpriteAnimation.Load()
	antagonist.JumpSpeed = player.Gravity
	players = append(players, antagonist)

	cam := NewCamera(player)
	defer func() {
		rl.CloseWindow()
		player.SpriteAnimation.Unload()
	}()
	rl.SetTargetFPS(60)

	layer := gameMap.Layers[0]
	clouds := InitClouds(gameMap)
	trees := InitTrees()

	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()

		player.Update(delta, envObjs)
		player.SpriteAnimation.Update(delta, &player)

		spriteCurrentState := player.SpriteAnimation.CurrentState
		currentFrame := player.SpriteAnimation.CurrentFrame
		currentTex := player.SpriteAnimation.StateTextures[spriteCurrentState]

		cam.Update(player, gameMap, delta)

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
					X:      float32(tileX) * tileWidth,
					Y:      float32(tileY) * tileHeight,
					Width:  tileWidth,
					Height: tileHeight,
				}

				destination := rl.Rectangle{
					X: float32(x) * tileWidth,
					Y: float32(y) * tileHeight,
					//scaling this piece of shit to get rid of the flickering
					Width:  tileWidth + 0.1,
					Height: tileHeight + 0.1,
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
		rl.DrawRectangleLinesEx(player.Box, 1, rl.Red)
		// for _, eo := range envObjs {
		// 	rl.DrawRectangleLinesEx(
		// 		eo.Rect,
		// 		1,
		// 		eo.Color,
		// 	)
		// }
		rl.EndMode2D()
		rl.DrawFPS(20, 20)

		rl.EndDrawing()
	}
}
