package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	StateIdle   = "idle"
	StateWalk   = "walk"
	StateRun    = "run"
	StateJump   = "jump"
	StateDash   = "dash"
	StateCombo1 = "combo1"
	StateCombo2 = "combo2"
	StateCombo3 = "combo3"
)

type Player struct {
	Box               rl.Rectangle
	FootPosition      rl.Vector2
	JumpSpeed         float32
	FacingRight       bool
	HorSpeed          float32
	SpriteAnimation   *SpriteAnimation
	Grounded          bool
	Gravity           float32
	DragSpeed         float32
	ComboCounter      int
	ComboResetCounter float32
	NextComboQueued   bool
}

func (p *Player) startCombo() {
	p.ComboCounter += 1
	if p.ComboCounter > 3 {
		p.ComboCounter = 1
		p.ComboResetCounter = 0
	}
	if p.ComboResetCounter <= 0 {
		p.ComboCounter = 1
		p.ComboResetCounter += 5
	}
	if p.ComboResetCounter < 3 {
		p.ComboResetCounter += 5
	}
	p.SpriteAnimation.CurrentFrame = 0
	p.SpriteAnimation.CurrentState = "combo" + strconv.Itoa(p.ComboCounter)
	p.SpriteAnimation.OneShot = true
}

func (p *Player) isInActionState() bool {
	cs := p.SpriteAnimation.CurrentState
	return cs == StateDash || cs == StateCombo1 || cs == StateCombo2 || cs == StateCombo3
}

func (p *Player) isInCombo() bool {
	cs := p.SpriteAnimation.CurrentState
	return cs == StateCombo1 || cs == StateCombo2 || cs == StateCombo3
}

func (p *Player) Update(
	delta float32,
	envObjs []EnvironmentObject,
) {
	p.ComboResetCounter -= 8 * delta
	if p.ComboResetCounter <= 0 {
		p.ComboCounter = 0
		p.ComboResetCounter = 0
	}
	// rl.TraceLog(rl.LogInfo, "ComboResetCounter: %v", p.ComboResetCounter)
	if p.JumpSpeed != 0 {
		p.SpriteAnimation.FramePinned = false
		p.Grounded = false
		if !p.isInActionState() {
			p.SpriteAnimation.CurrentState = StateJump
		}
	}
	if p.SpriteAnimation.CurrentState != StateDash {
		if rl.IsKeyDown(rl.KeyL) {
			p.SpriteAnimation.OneShot = true
			p.SpriteAnimation.CurrentFrame = 0
			p.SpriteAnimation.CurrentState = StateDash
			if p.FacingRight {
				p.HorSpeed = hor_speed * 3.5
			} else {
				p.HorSpeed = -hor_speed * 3.5
			}
		}
		if rl.IsKeyPressed(rl.KeyJ) {
			if p.isInCombo() {
				p.NextComboQueued = true
			} else {
				p.startCombo()
			}
		}
		if rl.IsKeyDown(rl.KeyD) &&
			rl.IsKeyDown(rl.KeyLeftShift) &&
			!rl.IsKeyDown(rl.KeyL) &&
			!p.isInCombo() {
			p.FacingRight = true
			p.HorSpeed = hor_speed * 2
			if p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = StateRun
			}
		} else if rl.IsKeyDown(rl.KeyD) &&
			!rl.IsKeyDown(rl.KeyL) &&
			!p.isInCombo() {
			p.FacingRight = true
			p.HorSpeed = hor_speed
			if p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = StateWalk
			}
		}
		if rl.IsKeyDown(rl.KeyA) &&
			rl.IsKeyDown(rl.KeyLeftShift) &&
			!rl.IsKeyDown(rl.KeyL) &&
			!p.isInCombo() {
			p.FacingRight = false
			p.HorSpeed = -hor_speed * 2
			if p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = StateRun
			}
		} else if rl.IsKeyDown(rl.KeyA) &&
			!rl.IsKeyDown(rl.KeyL) &&
			!p.isInCombo() {
			p.FacingRight = false
			p.HorSpeed = -hor_speed
			if p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = StateWalk
			}
		}
	}
	if rl.IsKeyPressed(rl.KeySpace) && p.Grounded {
		p.SpriteAnimation.CurrentState = StateJump
		if p.SpriteAnimation.CurrentState != StateDash {
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
		!p.isInActionState() {
		p.SpriteAnimation.FramePinned = true
		if p.JumpSpeed > 0 {
			p.SpriteAnimation.CurrentFrame = 4
		} else if p.JumpSpeed < 0 {
			p.SpriteAnimation.CurrentFrame = 2
		}
	}

	if !p.isInActionState() {
		if p.Grounded {
			if p.HorSpeed == 0 && p.JumpSpeed == 0 {
				p.SpriteAnimation.CurrentState = StateIdle
			} else if !rl.IsKeyDown(rl.KeyA) && !rl.IsKeyDown(rl.KeyD) && p.HorSpeed != 0 {
				p.SpriteAnimation.CurrentState = StateWalk
			}
		}
	}
	oldFootY := p.FootPosition.Y

	p.JumpSpeed += p.Gravity * delta
	p.FootPosition.Y += p.JumpSpeed * delta

	p.OnGrounded(envObjs, oldFootY)

	p.SpriteAnimation.Rect.X = p.FootPosition.X - p.SpriteAnimation.Rect.Width/2
	p.SpriteAnimation.Rect.Y = p.FootPosition.Y - p.SpriteAnimation.Rect.Height
	p.Box.X = p.FootPosition.X - 13
	p.Box.Y = p.FootPosition.Y + -60
}

func (p *Player) OnGrounded(
	envObjs []EnvironmentObject,
	oldFootPosY float32,
) {
	if p.HorSpeed == 0 && p.Grounded && !p.SpriteAnimation.OneShot {
		p.SpriteAnimation.CurrentState = StateIdle
	}
	for _, eo := range envObjs {
		if p.JumpSpeed >= 0 &&

			p.FootPosition.X >= eo.Rect.X &&
			p.FootPosition.X <= eo.Rect.X+eo.Rect.Width &&

			//i dont understand this, its from raylib by examples, ill ponder on it later
			oldFootPosY <= eo.Rect.Y &&
			p.FootPosition.Y >= eo.Rect.Y {

			p.JumpSpeed = 0
			p.FootPosition.Y = eo.Rect.Y
			p.Grounded = true

			p.SpriteAnimation.FramePinned = false
			p.Gravity = 900
			p.SpriteAnimation.FrameDuration = 0.099
		}
	}
}
