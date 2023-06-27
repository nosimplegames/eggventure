package game

import (
	"github.com/nosimplegames/eggventure/engine"
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/events"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/physics"
)

type Grunt struct {
	GruntCharacter

	detector          *AIBoxDetector
	isFollowingTarget bool
}

func (grunt *Grunt) Update() {
	if !grunt.isFollowingTarget {
		grunt.SetState(engine.Idle)
		grunt.SetMovingDirection(engine.NoMoving)
	}

	grunt.isFollowingTarget = false
	grunt.Character.Update()

}

func (grunt *Grunt) followTarget(targetPosition math.Vector) {
	isTargetAtLeft := targetPosition.X < grunt.GetPosition().X

	grunt.SetState(engine.Walking)
	grunt.isFollowingTarget = true

	if isTargetAtLeft {
		grunt.SetMovingDirection(engine.MovingDirectionLeft)
	} else {
		grunt.SetMovingDirection(engine.MovingDirectionRight)
	}
}

type GruntFactory struct {
}

func (factory GruntFactory) Create() *Grunt {
	grunt := &Grunt{}
	GruntCharacterFactory{}.Init(&grunt.GruntCharacter)

	grunt.SetPosition(math.Vector{
		X: res.GameSize.X * 0.8,
		Y: res.GameSize.Y * 0.8,
	})

	factory.createAIBoxDetector(grunt)

	return grunt
}

func (factory GruntFactory) createAIBoxDetector(grunt *Grunt) {
	grunt.detector = AIBoxDetectorFactory{
		Size: res.GruntBoxDetector,
	}.Create()
	grunt.detector.SetPosition(grunt.GetSize().By(0.5))
	physics.AddCollisionable(grunt.detector)
	core.EntityAdder{
		Parent: grunt,
		Child:  grunt.detector,
	}.Add()

	grunt.detector.OnPlayerDetected.AddTCallback(func(data events.SignalData) {
		playerPosition := data.(math.Vector)
		grunt.followTarget(playerPosition)
	})
}
