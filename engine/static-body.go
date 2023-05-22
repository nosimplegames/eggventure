package engine

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/physics"
)

type StaticBody struct {
	Size          math.Vector
	Position      math.Vector
	CollisionMask string
}

func (body StaticBody) GetSize() math.Vector {
	return body.Size
}

func (body StaticBody) GetPosition() math.Vector {
	return body.Position
}

func (body StaticBody) GetCollisionMask() string {
	return body.CollisionMask
}

func (body StaticBody) CanCollide() bool {
	return true
}

func (body StaticBody) CanCollideWith(collisionMask string) bool {
	return collisionMask != body.CollisionMask
}

func (body StaticBody) OnCollision(collision physics.Collision) {
}

func (body StaticBody) IsAlive() bool {
	return true
}
