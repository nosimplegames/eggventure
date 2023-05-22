package engine

type CharacterFalling struct {
	FSMState[*Character]
}

func (state CharacterFalling) OnEnter(target *Character) {
	target.SetTexture(target.FallingTexture)
	target.Body.Fall()
}

type CharacterFallingFactory struct {
}

func (factory CharacterFallingFactory) Create() *CharacterFalling {
	state := &CharacterFalling{}
	state.Id = Falling
	state.SupportedStates = []FSMStateId{
		Landing,
	}

	return state
}
