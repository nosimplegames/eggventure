package engine

import (
	"github.com/nosimplegames/ns-framework/core"
)

type CharacterWalking struct {
	FSMState[*Character]

	animation        core.IAnimation
	runningAnimation core.IAnimation
}

func (state *CharacterWalking) OnEnter(target *Character) {
	state.runningAnimation = state.animation.Copy(target)
	core.AddAnimation(state.runningAnimation)
	target.SetMovingDirection(MovingDirectionLeft)
}

func (state CharacterWalking) OnExit(target *Character) {
	state.runningAnimation.Stop()
}

type CharacterWalkingFactory struct {
	Animation core.IAnimation
}

func (factory CharacterWalkingFactory) Create() *CharacterWalking {
	state := &CharacterWalking{}
	state.Id = Walking
	state.SupportedStates = []FSMStateId{
		Idle,
		Jumping,
	}
	state.animation = factory.Animation

	return state
}
