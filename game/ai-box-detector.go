package game

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/events"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/physics"
)

type AIBoxDetector struct {
	core.Entity

	OnPlayerDetected events.Signal
}

func (detector AIBoxDetector) GetCollisionMask() string {
	return "ai-box-detector"
}

func (detector AIBoxDetector) CanCollide() bool {
	return true
}

func (detector AIBoxDetector) CanCollideWith(collisionMask string) bool {
	return collisionMask == "player"
}

func (detector AIBoxDetector) OnCollision(collision physics.Collision) {
	playerPosition := collision.Another.GetPosition()
	detector.OnPlayerDetected.TFire(playerPosition)
}

type AIBoxDetectorFactory struct {
	Size math.Vector
}

func (factory AIBoxDetectorFactory) Create() *AIBoxDetector {
	detector := &AIBoxDetector{}
	detector.SetSize(factory.Size)

	return detector
}
