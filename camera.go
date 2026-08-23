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
		Zoom:     1,
	}
	return &Camera{
		Cam: cam,
	}
}

var edgePosX float32 = 0

func (c *Camera) Update(
	p Player,
	delta float32,

) {

	c.Cam.Target = p.FootPosition
}
