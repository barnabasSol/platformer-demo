package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteAnimation struct {
	CurrentState  string
	StateTextures map[string][]rl.Texture2D
	CurrentFrame  int
	FrameDuration float32
	Timer         float32
	Src           rl.Rectangle
	Dst           rl.Rectangle
}

func (s *SpriteAnimation) Update(
	delta float32,
	player *Player,
) {
	s.Timer += delta
	var srcRectWidth float32
	if player.FacingRight {
		srcRectWidth = float32(s.StateTextures[s.CurrentState][s.CurrentFrame].Width)
	} else {
		srcRectWidth = -float32(s.StateTextures[s.CurrentState][s.CurrentFrame].Width)
	}
	src := rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  srcRectWidth,
		Height: float32(s.StateTextures[s.CurrentState][s.CurrentFrame].Height),
	}

	s.Src = src

	dst := player.Rect
	if s.Timer >= s.FrameDuration {
		s.CurrentFrame = (s.CurrentFrame + 1) % len(s.StateTextures[s.CurrentState])
		s.Timer = 0
	}
	dst.Y += 66
	s.Dst = dst

}

func (s *SpriteAnimation) Load() {
	states := []string{"idle", "walk", "run", "jump"}
	for _, state := range states {
		paths := getTextures(state)
		var tex []rl.Texture2D
		for _, v := range paths {
			tex = append(tex, rl.LoadTexture(v))
		}
		s.StateTextures[state] = tex
	}
}
