package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
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
		Zoom:     1.1,
	}
	return &Camera{
		Cam: cam,
	}
}

var vecX = 60
var edgePosX float32 = 0

func (c *Camera) Update(
	p Player,
	delta float32,
	gm Map,

) {
	if c.EdgeReached {
		c.Cam.Target.X = edgePosX
		c.Cam.Target.Y = p.FootPosition.Y
	} else {
		c.Cam.Target = p.FootPosition
	}
	if !c.EdgeReached &&
		c.Cam.Target.X <= float32(gm.MinWidth) ||
		c.Cam.Target.X+c.Cam.Offset.X >= float32(gm.MaxWidth) {
		c.EdgeReached = true
		edgePosX = c.Cam.Target.X
	}
	//right
	if c.EdgeReached &&
		p.FootPosition.X > (float32(gm.MinWidth)+float32(gm.MaxWidth))/2 &&
		p.FootPosition.X < edgePosX {
		c.EdgeReached = false
	}
	//left
	if c.EdgeReached &&
		p.FootPosition.X < (float32(gm.MinWidth)+float32(gm.MaxWidth))/2 &&
		p.FootPosition.X > edgePosX {
		c.EdgeReached = false
	}
}
