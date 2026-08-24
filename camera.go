package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

type Camera struct {
	Cam         rl.Camera2D
	EdgeReached bool
}

func NewCamera(p Player) *Camera {
	cam := rl.Camera2D{
		Target: p.FootPosition,
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

var edgePosX float32 = 0

func (c *Camera) Update(
	p Player,
	gameMap *tiled.Map,
	delta float32,

) {

	//the 10 added/subtracted is to eliminate the extra view that goes beyond the map,
	//shitty solution, ill get to do something better eventually
	left := c.Cam.Target.X - c.Cam.Offset.X - 10/c.Cam.Zoom

	right := c.Cam.Target.X + (float32(width)-c.Cam.Offset.X+10)/c.Cam.Zoom

	// top := c.Cam.Target.Y - c.Cam.Offset.Y/c.Cam.Zoom

	// bottom := c.Cam.Target.Y + (float32(height)-c.Cam.Offset.Y)/c.Cam.Zoom

	gameFullWidth := gameMap.Width * gameMap.TileWidth
	// gameFullHeight := gameMap.Height * gameMap.TileHeight

	center := float32(gameMap.Layers[0].OffsetX) + float32(gameFullWidth)/2
	if !c.EdgeReached {
		if left <= float32(gameMap.Layers[0].OffsetX) ||
			right >= float32(gameMap.Layers[0].OffsetX)+float32(gameFullWidth) {
			edgePosX = p.FootPosition.X
			c.EdgeReached = true
		}
	}
	if !c.EdgeReached {
		c.Cam.Target.X = p.FootPosition.X
	} else {
		c.Cam.Target.X = edgePosX
	}

	if c.EdgeReached &&
		p.FootPosition.X < center &&
		p.FootPosition.X > edgePosX {
		c.EdgeReached = false
	}

	if c.EdgeReached &&
		p.FootPosition.X > center &&
		p.FootPosition.X < edgePosX {
		c.EdgeReached = false
	}

	c.Cam.Target.Y = p.FootPosition.Y + 140

}
