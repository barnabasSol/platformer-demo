package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

type Camera struct {
	Cam          rl.Camera2D
	EdgeReached  bool
	EdgeReachedY bool
}

func NewCamera(p Player) *Camera {
	cam := rl.Camera2D{
		Target: rl.NewVector2(p.Box.X, p.Box.Y),
		Offset: rl.Vector2{
			X: float32(width) / float32(2),
			Y: float32(height) / 1.3,
		},
		Rotation: 0,
		Zoom:     1,
	}
	return &Camera{
		Cam: cam,
	}
}

func (c *Camera) Update(
	p Player,
	gameMap *tiled.Map,
	delta float32,
) {
	mapLeft := float32(gameMap.Layers[0].OffsetX)
	mapTop := float32(gameMap.Layers[0].OffsetY)

	mapWidth := float32(gameMap.Width * gameMap.TileWidth)
	mapHeight := float32(gameMap.Height * gameMap.TileHeight)

	mapRight := mapLeft + mapWidth
	mapBottom := mapTop + mapHeight

	playerX := p.Box.X + p.Box.Width/2
	playerY := p.Box.Y + p.Box.Height/2

	targetX := playerX
	targetY := playerY + 40

	leftSpace := c.Cam.Offset.X / c.Cam.Zoom
	rightSpace := (float32(width) - c.Cam.Offset.X) / c.Cam.Zoom

	bottomSpace := (float32(height) - c.Cam.Offset.Y) / c.Cam.Zoom

	if targetX-leftSpace < mapLeft {
		targetX = mapLeft + leftSpace
	}

	if targetX+rightSpace > mapRight {
		targetX = mapRight - rightSpace
	}

	// if targetY-topSpace < mapTop {
	// 	targetY = mapTop + topSpace
	// }

	if targetY+bottomSpace > mapBottom {
		targetY = mapBottom - bottomSpace
	}

	c.Cam.Target.X = targetX
	c.Cam.Target.Y = targetY
}
