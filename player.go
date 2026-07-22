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
	Grounded        bool
	Gravity         float32
	DragSpeed       float32
}

func (p *Player) Update(
	delta float32,
	envItems []EnvironmentItem,
) {
	if rl.IsKeyDown(rl.KeyD) && rl.IsKeyDown(rl.KeyLeftShift) {
		p.FacingRight = true
		p.HorSpeed = hor_speed * 2
		if p.JumpSpeed == 0 {
			p.SpriteAnimation.CurrentState = "run"
		}
	} else if rl.IsKeyDown(rl.KeyD) {
		p.FacingRight = true
		p.HorSpeed = hor_speed
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
	if p.Grounded {
		if p.HorSpeed == 0 && p.JumpSpeed == 0 {
			p.SpriteAnimation.CurrentState = "idle"
		} else if !rl.IsKeyDown(rl.KeyD) && !rl.IsKeyDown(rl.KeyA) && p.HorSpeed != 0 {
			p.SpriteAnimation.CurrentState = "walk"
		}
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.Grounded {
		p.SpriteAnimation.CurrentFrame = 0
		p.SpriteAnimation.CurrentState = "jump"
		p.JumpSpeed = jump_speed
		p.Grounded = false
	}

	p.JumpSpeed += p.Gravity * delta
	p.Position.Y += p.JumpSpeed * delta
	p.SpriteAnimation.Rect.X = p.Position.X - p.SpriteAnimation.Rect.Width/2
	p.SpriteAnimation.Rect.Y = p.Position.Y - p.SpriteAnimation.Rect.Height

	p.OnGrounded(envItems)

	if !p.Grounded {
		if p.JumpSpeed > 0 {
			p.SpriteAnimation.CurrentFrame = 4
			p.SpriteAnimation.FramePinned = true
		} else if p.JumpSpeed < 0 {
			p.SpriteAnimation.CurrentFrame = 2
			p.SpriteAnimation.FramePinned = true
		}
	}
}

func (p *Player) OnGrounded(envItems []EnvironmentItem) {
	if p.Position.Y >= float32(floor) {
		p.JumpSpeed = 0
		p.Position.Y = float32(floor)
		p.Grounded = true
		p.SpriteAnimation.FramePinned = false
		p.Gravity = 900
		p.SpriteAnimation.FrameDuration = 0.099
	}
	for _, ei := range envItems {
		if p.Position.Y >= ei.Rect.Y &&
			p.Position.X >= float32(ei.Rect.X) &&
			p.Position.X <= ei.Rect.X+float32(ei.Rect.Width) &&
			p.Position.Y <= ei.Rect.Y+ei.Rect.Height &&
			p.JumpSpeed >= 0 {

			p.JumpSpeed = 0
			p.Position.Y = ei.Rect.Y
			p.Grounded = true
			p.SpriteAnimation.FramePinned = false
			p.Gravity = 900
			p.SpriteAnimation.FrameDuration = 0.099
		}
	}
}
