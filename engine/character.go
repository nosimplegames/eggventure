package engine

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/physics"
	"github.com/nosimplegames/ns-framework/render"
)

type MovingDirection int

const (
	NoMoving MovingDirection = iota
	MovingDirectionLeft
	MovingDirectionRight
)

type Character struct {
	entities.Sprite

	FSM           IFSM[*Character]
	Body          DynamicBody
	CollisionMask string

	IdleTexture    render.Texture
	FallingTexture render.Texture
	JumpingTexture render.Texture
}

func (character *Character) SetState(state CharacterState) {
	character.FSM.SetState(state)
}

func (character *Character) SetMovingDirection(direction MovingDirection) {
	character.Body.SetMovingDirection(direction)

	switch direction {
	case MovingDirectionLeft:
		character.SetScale(math.Vector{
			X: -1,
			Y: 1,
		})
	case MovingDirectionRight:
		character.SetScale(math.Vector{
			X: 1,
			Y: 1,
		})
	}
}

func (character Character) GetMovingDirection() MovingDirection {
	return character.Body.MovingDirection
}

func (character *Character) Update() {
	character.Body.Update()
	character.FSM.Update()

	character.Move(character.Body.Speed)
}

func (character Character) GetCollisionMask() string {
	return character.CollisionMask
}

func (character Character) CanCollide() bool {
	return true
}

func (character Character) CanCollideWith(collisionMask string) bool {
	return collisionMask == "floor" ||
		collisionMask == "wall"
}

func (character *Character) OnCollision(collision physics.Collision) {
	switch collision.AnotherCollisionMask {
	case "floor":
		character.OnFloorCollision(collision)
	case "wall":
		character.OnWallCollision(collision)
	}
}

func (character *Character) OnFloorCollision(collision physics.Collision) {
	characterPosition := character.GetPosition()
	yResolution := collision.CollisionResolverCalculator.CalculateYResolution()
	characterPosition.Y += yResolution

	character.SetPosition(characterPosition)
	character.FSM.SetState(Landing)
}

func (character *Character) OnWallCollision(collision physics.Collision) {
	characterPosition := character.GetPosition()
	xResolution := collision.CollisionResolverCalculator.CalculateXResolution()
	characterPosition.X += xResolution

	character.SetPosition(characterPosition)
}

func (character *Character) IsAlive() bool {
	return true
}

type CharacterFactory struct {
	IdleTexture      render.Texture
	FallingTexture   render.Texture
	JumpingTexture   render.Texture
	WalkingAnimation core.IAnimation
}

func (factory CharacterFactory) Init(character *Character) {
	character.IdleTexture = factory.IdleTexture
	character.FallingTexture = factory.FallingTexture
	character.JumpingTexture = factory.JumpingTexture

	character.FSM = &FSM[*Character]{
		Target: character,
	}
	character.FSM.AddState(CharacterFallingFactory{}.Create())
	character.FSM.AddState(CharacterIdleFactory{}.Create())
	character.FSM.AddState(CharacterLandingFactory{}.Create())
	character.FSM.AddState(CharacterWalkingFactory{
		Animation: factory.WalkingAnimation,
	}.Create())
	character.FSM.AddState(CharacterJumpingFactory{}.Create())
}
