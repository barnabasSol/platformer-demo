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
) {
	if c.EdgeReached {
		println("eat me")
		c.Cam.Target.X = edgePosX
	} else {
		c.Cam.Target = p.FootPosition
	}
	if !c.EdgeReached && c.Cam.Target.X <= 370 || c.Cam.Target.X+c.Cam.Offset.X >= 1300 {
		c.EdgeReached = true
		edgePosX = c.Cam.Target.X
	}
	if c.EdgeReached &&
		p.FootPosition.X > (370+1300)/2 &&
		p.FootPosition.X < edgePosX {
		c.EdgeReached = false
	}
	if c.EdgeReached &&
		p.FootPosition.X < (370+1300)/2 &&
		p.FootPosition.X > edgePosX {
		c.EdgeReached = false
	}
}
