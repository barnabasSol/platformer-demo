package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position        rl.Vector2
	JumpSpeed       float32
	FacingRight     bool
	HorSpeed        float32
	SpriteAnimation *SpriteAnimation
	CanJump         bool
	Gravity         float32
	DragSpeed       float32
}

func (p *Player) Update(delta float32) {
	if rl.IsKeyDown(rl.KeyD) && rl.IsKeyDown(rl.KeyLeftShift) {
		p.FacingRight = true
		p.HorSpeed = hor_speed * 2
		if p.JumpSpeed == 0 {
			p.SpriteAnimation.CurrentState = "run"
		}
	} else if rl.IsKeyDown(rl.KeyD) && p.JumpSpeed == 0 {
		p.FacingRight = true
		p.HorSpeed = hor_speed
		if p.JumpSpeed == 0 {
			p.SpriteAnimation.CurrentState = "walk"
		}
	} else if p.CanJump {
		if p.JumpSpeed == 0 {
			p.SpriteAnimation.CurrentState = "walk"
		}
	}
	if rl.IsKeyDown(rl.KeyA) && rl.IsKeyDown(rl.KeyLeftShift) {
		p.FacingRight = false
		p.HorSpeed = -hor_speed * 2
		if p.JumpSpeed == 0 {
			p.SpriteAnimation.CurrentState = "run"
		}
	} else if rl.IsKeyDown(rl.KeyA) {
		p.FacingRight = false
		p.HorSpeed = -hor_speed
		if p.JumpSpeed == 0 {
			p.SpriteAnimation.CurrentState = "walk"
		}
	}
	if p.FacingRight {
		p.Position.X += p.HorSpeed * delta
		p.HorSpeed += p.DragSpeed * delta
		if p.HorSpeed <= 0 {
			p.HorSpeed = 0
		}
	} else {
		p.Position.X += p.HorSpeed * delta
		p.HorSpeed += -1 * p.DragSpeed * delta
		if p.HorSpeed >= 0 {
			p.HorSpeed = 0
		}
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.CanJump {
		p.SpriteAnimation.CurrentFrame = 0
		p.SpriteAnimation.CurrentState = "jump"
		p.JumpSpeed = jump_speed
		p.CanJump = false
	}

	p.JumpSpeed += p.Gravity * delta
	p.Position.Y += p.JumpSpeed * delta
	p.SpriteAnimation.Rect.X = p.Position.X - p.SpriteAnimation.Rect.Width/2
	p.SpriteAnimation.Rect.Y = p.Position.Y - p.SpriteAnimation.Rect.Height

	if p.Position.Y >= float32(floor) {
		p.Position.Y = float32(floor)
		p.Gravity = 900
		p.JumpSpeed = 0
		p.CanJump = true
		if p.HorSpeed == 0 {
			p.SpriteAnimation.CurrentState = "idle"
		}
	}
}
