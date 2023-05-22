package engine

type CharacterLanding struct {
	FSMState[*Character]
}

func (state CharacterLanding) OnEnter(target *Character) {
	target.Body.StopFalling()
}

func (state CharacterLanding) Update(target *Character) {
	target.SetState(Idle)
}

type CharacterLandingFactory struct {
}

func (factory CharacterLandingFactory) Create() *CharacterLanding {
	state := &CharacterLanding{}
	state.Id = Landing
	state.SupportedStates = []FSMStateId{
		Idle,
	}

	return state
}
