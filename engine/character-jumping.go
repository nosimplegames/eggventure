package engine

type CharacterJumping struct {
	FSMState[*Character]
}

func (state CharacterJumping) OnEnter(target *Character) {
	target.SetTexture(target.JumpingTexture)
	target.Body.Jump()
}

func (state CharacterJumping) Update(target *Character) {
	if target.Body.HasStopJumping() {
		target.SetState(Falling)
	}
}

type CharacterJumpingFactory struct {
}

func (factory CharacterJumpingFactory) Create() *CharacterJumping {
	state := &CharacterJumping{}
	state.Id = Jumping
	state.SupportedStates = []FSMStateId{
		Falling,
	}

	return state
}
