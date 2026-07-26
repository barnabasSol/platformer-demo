package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteAnimation struct {
	CurrentState  string
	StateTextures map[string][]rl.Texture2D
	CurrentFrame  int
	FramePinned   bool
	FrameDuration float32
	Timer         float32
	Rect          rl.Rectangle
	Src           rl.Rectangle
	Dst           rl.Rectangle
	OneShot       bool
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

	dst := s.Rect
	if !s.FramePinned && s.Timer >= s.FrameDuration {
		if s.OneShot &&
			s.CurrentFrame == len(s.StateTextures[s.CurrentState])-1 {
			s.OneShot = false
			if player.NextComboQueued {
				player.NextComboQueued = false
				player.startCombo()
			} else {
				s.CurrentState = StateIdle
			}
		} else {
			s.CurrentFrame = (s.CurrentFrame + 1) % len(s.StateTextures[s.CurrentState])
			s.Timer = 0
		}
	}
	dst.Y += 66
	s.Dst = dst
}

func (s *SpriteAnimation) Load() {
	states := []string{
		StateIdle,
		StateWalk,
		StateRun,
		StateJump,
		StateDash,
		StateCombo1,
		StateCombo2,
		StateCombo3,
	}
	for _, state := range states {
		paths := getTextures(state)
		var tex []rl.Texture2D
		for _, v := range paths {
			tex = append(tex, rl.LoadTexture(v))
		}
		s.StateTextures[state] = tex
	}
}

func (s *SpriteAnimation) Unload() {
	states := []string{
		StateIdle,
		StateWalk,
		StateRun,
		StateJump,
		StateDash,
		StateCombo1,
		StateCombo2,
		StateCombo3,
	}
	for _, state := range states {
		for _, tex := range s.StateTextures[state] {
			rl.UnloadTexture(tex)
		}
	}
}
