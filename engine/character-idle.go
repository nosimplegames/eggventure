package engine

type CharacterIdle struct {
	FSMState[*Character]
}

func (state CharacterIdle) OnEnter(target *Character) {
	target.SetTexture(target.IdleTexture)
	target.SetMovingDirection(NoMoving)
}

type CharacterIdleFactory struct {
}

func (factory CharacterIdleFactory) Create() *CharacterIdle {
	state := &CharacterIdle{}
	state.Id = Idle
	state.SupportedStates = []FSMStateId{
		Jumping,
		Falling,
		Walking,
	}

	return state
}
