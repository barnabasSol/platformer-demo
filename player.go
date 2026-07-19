package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position        rl.Vector2
	Speed           float32
	FacingRight     bool
	HorSpeed        float32
	SpriteAnimation *SpriteAnimation
	CanJump         bool
}

func (p *Player) Update(delta float32) {
	if rl.IsKeyDown(rl.KeyD) && rl.IsKeyDown(rl.KeyLeftShift) {
		p.FacingRight = true
		p.HorSpeed = hor_speed * 2
		p.SpriteAnimation.CurrentState = "run"
	} else if rl.IsKeyDown(rl.KeyD) {
		p.FacingRight = true
		p.HorSpeed = hor_speed
		p.SpriteAnimation.CurrentState = "walk"
	} else {
		p.SpriteAnimation.CurrentState = "walk"
	}
	if rl.IsKeyDown(rl.KeyA) && rl.IsKeyDown(rl.KeyLeftShift) {
		p.SpriteAnimation.CurrentState = "run"
		p.FacingRight = false
		p.HorSpeed = -hor_speed * 2
	} else if rl.IsKeyDown(rl.KeyA) {
		p.SpriteAnimation.CurrentState = "walk"
		p.FacingRight = false
		p.HorSpeed = -hor_speed
	}
	if p.FacingRight {
		p.Position.X += p.HorSpeed * delta
		p.HorSpeed += drag_speed * delta
		if p.HorSpeed <= 0 {
			p.HorSpeed = 0
			p.SpriteAnimation.CurrentState = "idle"
		}
	} else {
		p.Position.X += p.HorSpeed * delta
		p.HorSpeed += -1 * drag_speed * delta
		if p.HorSpeed >= 0 {
			p.HorSpeed = 0
			p.SpriteAnimation.CurrentState = "idle"
		}
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.CanJump {
		p.Speed = jump_speed
		p.CanJump = false
	}

	p.Speed += gravity * delta
	p.Position.Y += p.Speed * delta
	p.SpriteAnimation.Rect.X = p.Position.X - p.SpriteAnimation.Rect.Width/2
	p.SpriteAnimation.Rect.Y = p.Position.Y - p.SpriteAnimation.Rect.Height

	if p.Position.Y >= float32(floor) {
		p.Position.Y = float32(floor)
		gravity = 900
		p.Speed = 0
		p.CanJump = true
	}
}
