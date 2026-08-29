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
		p.HorSpeed += p.DragSpeed * delta
		if p.HorSpeed <= 0 {
			p.HorSpeed = 0
		}
	} else {
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
	oldX := p.Box.X
	oldY := p.Box.Y

	p.JumpSpeed += p.Gravity * delta
	// p.FootPosition.Y += p.JumpSpeed * delta
	// p.FootPosition.X += p.HorSpeed * delta

	p.Box.X += p.HorSpeed * delta
	p.Box.Y += p.JumpSpeed * delta

	p.OnContact(envObjs, oldX, oldY)

	//migrated the foot dot to the box
	p.SpriteAnimation.Rect.X = p.Box.X + p.Box.Width/2 - p.SpriteAnimation.Rect.Width/2
	p.SpriteAnimation.Rect.Y = p.Box.Y + p.Box.Height + -p.SpriteAnimation.Rect.Height
}

var debugInt int

func (p *Player) OnContact(
	envObjs []EnvironmentObject,
	oldX float32,
	oldY float32,
) {
	p.Grounded = false

	for _, eo := range envObjs {
		if !rl.CheckCollisionRecs(p.Box, eo.Rect) {
			continue
		}

		oldBottom := oldY + p.Box.Height
		oldTop := oldY
		oldRight := oldX + p.Box.Width
		oldLeft := oldX

		objectTop := eo.Rect.Y
		objectBottom := eo.Rect.Y + eo.Rect.Height
		objectLeft := eo.Rect.X
		objectRight := eo.Rect.X + eo.Rect.Width

		// Falling onto platform
		if p.JumpSpeed >= 0 &&
			oldBottom <= objectTop {

			p.Box.Y = objectTop - p.Box.Height
			p.Gravity = 900
			p.JumpSpeed = 0
			p.Grounded = true
			p.SpriteAnimation.FramePinned = false

			// Jumping into underside
		} else if p.JumpSpeed < 0 &&
			oldTop >= objectBottom {

			p.Box.Y = objectBottom
			p.JumpSpeed = 0

			// Moving right into wall
		} else if p.HorSpeed > 0 &&
			oldRight <= objectLeft {

			p.Box.X = objectLeft - p.Box.Width
			p.HorSpeed = 0

			// Moving left into wall
		} else if p.HorSpeed < 0 &&
			oldLeft >= objectRight {

			p.Box.X = objectRight
			p.HorSpeed = 0
		}
	}
}
