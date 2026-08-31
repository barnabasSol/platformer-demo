package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Antagonist struct {
	Box             rl.Rectangle
	JumpSpeed       float32
	FacingRight     bool
	SpriteAnimation *SpriteAnimation
	Grounded        bool
	Gravity         float32
}

func (p *Antagonist) Update(
	delta float32,
	envObjs []EnvironmentObject,
) {
	// rl.TraceLog(rl.LogInfo, "ComboResetCounter: %v", p.ComboResetCounter)
	if p.JumpSpeed != 0 {
		p.SpriteAnimation.FramePinned = false
		p.Grounded = false
		p.SpriteAnimation.CurrentState = StateJump
	}

	oldX := p.Box.X
	oldY := p.Box.Y

	p.JumpSpeed += p.Gravity * delta
	// p.FootPosition.Y += p.JumpSpeed * delta
	// p.FootPosition.X += p.HorSpeed * delta

	p.Box.Y += p.JumpSpeed * delta
	if !p.Grounded {
		p.SpriteAnimation.FramePinned = true
		if p.JumpSpeed > 0 {
			p.SpriteAnimation.CurrentFrame = 4
		} else if p.JumpSpeed < 0 {
			p.SpriteAnimation.CurrentFrame = 2
		}
	}

	p.OnContact(envObjs, oldX, oldY)

	//migrated the foot dot to the box
	p.SpriteAnimation.Rect.X = p.Box.X + p.Box.Width/2 - p.SpriteAnimation.Rect.Width/2
	p.SpriteAnimation.Rect.Y = p.Box.Y + p.Box.Height + -p.SpriteAnimation.Rect.Height
}

func (p *Antagonist) OnContact(
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

		if oldBottom <= objectTop {
			p.SpriteAnimation.CurrentState = StateIdle
			p.Box.Y = objectTop - p.Box.Height
			p.Gravity = 900
			p.JumpSpeed = 0
			p.Grounded = true
			p.SpriteAnimation.FramePinned = false

		} else if oldTop >= objectBottom {

			p.Box.Y = objectBottom
			p.JumpSpeed = 0

		} else if oldRight <= objectLeft {

			p.Box.X = objectLeft - p.Box.Width

		} else if oldLeft >= objectRight {

			p.Box.X = objectRight
		}
	}
}

func (p Antagonist) Draw() {
	spriteCurrentState := p.SpriteAnimation.CurrentState
	currentFrame := p.SpriteAnimation.CurrentFrame
	currentTex := p.SpriteAnimation.StateTextures[spriteCurrentState]
	rl.DrawTexturePro(
		currentTex[currentFrame],
		p.SpriteAnimation.Src,
		p.SpriteAnimation.Dst,
		rl.Vector2{},
		0,
		rl.White,
	)

}
