package main

import rl "github.com/gen2brain/raylib-go/raylib"

type EnvironmentObject struct {
	Rect     rl.Rectangle
	Color    rl.Color
	Blocking bool
}
