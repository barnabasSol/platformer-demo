package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

type Cloud struct {
	X       int
	Y       int
	Texture rl.Texture2D
}

func InitClouds(gameMap *tiled.Map) []Cloud {
	var clouds []Cloud
	tex := rl.LoadTexture("res/day_platformer/PNG/clouds.png")
	for _, il := range gameMap.ImageLayers {
		if il.Name == "Cloud" {
			clouds = append(clouds, Cloud{
				X:       int(il.OffsetX),
				Y:       int(il.OffsetY),
				Texture: tex,
			})
		}
	}
	return clouds
}

func UpdateClouds() {

}

func DrawClouds() {

}

type Tree struct {
	Texture rl.Texture2D
}

func InitTrees() Tree {
	tex := rl.LoadTexture("res/day_platformer/PNG/trees.png")
	return Tree{
		Texture: tex,
	}
}

func (t Tree) DrawTrees(gameMap *tiled.Map) {
	for x := int32(1); x < int32(gameMap.Width*gameMap.TileWidth); x += t.Texture.Width {
		rl.DrawTexture(
			t.Texture,
			x,
			93,
			rl.White,
		)
	}
}
