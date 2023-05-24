package engine

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/utils"
)

type DynamicBodyState int

const (
	DynamicBodyFalling DynamicBodyState = iota
	DynamicBodyOnFloor
	DynamicBodyJumping
)

type DynamicBody struct {
	Speed           math.Vector
	State           DynamicBodyState
	MovingDirection MovingDirection
}

func (body *DynamicBody) Update() {
	switch body.State {
	case DynamicBodyFalling:
		body.Speed.Y += 9.8 * utils.FrameTime
	case DynamicBodyJumping:
		body.Speed.Y += 9.8 * utils.FrameTime
	}

	switch body.MovingDirection {
	case NoMoving:
		body.Speed.X = 0
	case MovingDirectionLeft:
		body.Speed.X = -69.9 * utils.FrameTime
	case MovingDirectionRight:
		body.Speed.X = 69.9 * utils.FrameTime
	}
}

func (body *DynamicBody) StopFalling() {
	body.State = DynamicBodyOnFloor
	body.Speed.Y = 0
}

func (body *DynamicBody) Fall() {
	body.State = DynamicBodyFalling
	body.Speed.Y = 0
}

func (body *DynamicBody) Jump() {
	body.State = DynamicBodyJumping
	body.Speed.Y -= 4
}

func (body DynamicBody) HasStopJumping() bool {
	isJumping := body.State == DynamicBodyJumping

	if !isJumping {
		return true
	}

	isJumping = body.Speed.Y >= 0.0

	return isJumping
}

func (body *DynamicBody) SetMovingDirection(movingDirection MovingDirection) {
	body.MovingDirection = movingDirection
}
