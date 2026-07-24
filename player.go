package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	FootPosition    rl.Vector2
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
	if p.JumpSpeed > 0 {
		p.Grounded = false
		if p.SpriteAnimation.CurrentState != "dash" {
			p.SpriteAnimation.CurrentState = "jump"
		}
	}
	if p.SpriteAnimation.CurrentState != "dash" {
		if rl.IsKeyDown(rl.KeyL) {
			p.SpriteAnimation.OneShot = true
			p.SpriteAnimation.CurrentFrame = 0
			p.SpriteAnimation.CurrentState = "dash"
			if p.FacingRight {
				p.HorSpeed = hor_speed * 3.5
			} else {
				p.HorSpeed = -hor_speed * 3.5
			}
		}

		if rl.IsKeyDown(rl.KeyJ) {
			p.SpriteAnimation.OneShot = true
			p.SpriteAnimation.CurrentState = "combo1"
			p.SpriteAnimation.CurrentFrame = 0
		}
		if rl.IsKeyDown(rl.KeyD) && rl.IsKeyDown(rl.KeyLeftShift) && !rl.IsKeyDown(rl.KeyL) {
			p.FacingRight = true
			p.HorSpeed = hor_speed * 2
			if p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = "run"
			}
		} else if rl.IsKeyDown(rl.KeyD) && !rl.IsKeyDown(rl.KeyL) {
			p.FacingRight = true
			p.HorSpeed = hor_speed
			if p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = "walk"
			}
		}
		if rl.IsKeyDown(rl.KeyA) && rl.IsKeyDown(rl.KeyLeftShift) && !rl.IsKeyDown(rl.KeyL) {
			p.FacingRight = false
			p.HorSpeed = -hor_speed * 2
			if p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = "run"
			}
		} else if rl.IsKeyDown(rl.KeyA) && !rl.IsKeyDown(rl.KeyL) {
			p.FacingRight = false
			p.HorSpeed = -hor_speed
			if p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = "walk"
			}
		}
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.Grounded {
		p.SpriteAnimation.CurrentState = "jump"
		if p.SpriteAnimation.CurrentState != "dash" {
			p.JumpSpeed = jump_speed
		}
		p.Grounded = false
	}
	if p.FacingRight {
		p.FootPosition.X += p.HorSpeed * delta
		p.HorSpeed += p.DragSpeed * delta
		if p.HorSpeed <= 0 {
			p.HorSpeed = 0
		}
	} else {
		p.FootPosition.X += p.HorSpeed * delta
		p.HorSpeed += -1 * p.DragSpeed * delta
		if p.HorSpeed >= 0 {
			p.HorSpeed = 0
		}
	}

	if !p.Grounded &&
		p.SpriteAnimation.CurrentState != "dash" &&
		p.SpriteAnimation.CurrentState != "combo1" {
		p.SpriteAnimation.FramePinned = true
		if p.JumpSpeed > 0 {
			p.SpriteAnimation.CurrentFrame = 4
		} else if p.JumpSpeed < 0 {
			p.SpriteAnimation.CurrentFrame = 2
		}
	}

	if p.SpriteAnimation.CurrentState != "dash" &&
		p.SpriteAnimation.CurrentState != "combo1" {
		if p.Grounded {
			if p.HorSpeed == 0 && p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = "idle"
			} else if !rl.IsKeyDown(rl.KeyA) && !rl.IsKeyDown(rl.KeyD) && p.HorSpeed != 0 {
				p.SpriteAnimation.CurrentState = "walk"
			}
		}
	}

	p.JumpSpeed += p.Gravity * delta
	p.FootPosition.Y += p.JumpSpeed * delta

	p.OnGrounded(envItems)

	p.SpriteAnimation.Rect.X = p.FootPosition.X - p.SpriteAnimation.Rect.Width/2
	p.SpriteAnimation.Rect.Y = p.FootPosition.Y - p.SpriteAnimation.Rect.Height
}

func (p *Player) OnGrounded(envItems []EnvironmentItem) {
	if p.HorSpeed == 0 && p.Grounded && !p.SpriteAnimation.OneShot {
		p.SpriteAnimation.CurrentState = "idle"
	}
	if p.FootPosition.Y >= float32(floor) {

		p.JumpSpeed = 0
		p.FootPosition.Y = float32(floor)
		p.Grounded = true
		p.SpriteAnimation.FramePinned = false
		p.Gravity = 900
		p.SpriteAnimation.FrameDuration = 0.099

	}
	for _, ei := range envItems {
		if p.FootPosition.Y >= ei.Rect.Y &&
			p.FootPosition.X >= float32(ei.Rect.X) &&
			p.FootPosition.X <= ei.Rect.X+float32(ei.Rect.Width) &&
			p.FootPosition.Y <= ei.Rect.Y+ei.Rect.Height &&
			p.JumpSpeed >= 0 {

			p.JumpSpeed = 0
			p.FootPosition.Y = ei.Rect.Y
			p.Grounded = true
			p.SpriteAnimation.FramePinned = false
			p.Gravity = 900
			p.SpriteAnimation.FrameDuration = 0.099

		}
	}
}
