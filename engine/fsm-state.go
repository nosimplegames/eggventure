package engine

type FSMState[T any] struct {
	Id              FSMStateId
	SupportedStates []FSMStateId
}

func (state FSMState[T]) GetId() FSMStateId {
	return state.Id
}

func (state FSMState[T]) CanMoveTo(stateId FSMStateId) bool {
	for _, supportedState := range state.SupportedStates {
		isSupportedState := supportedState == stateId

		if isSupportedState {
			return true
		}
	}

	return false
}

func (state FSMState[T]) OnEnter(target T) {
}

func (state FSMState[T]) Update(target T) {
}

func (state FSMState[T]) OnExit(target T) {
}
