package engine

import "fmt"

type FSM[T any] struct {
	Target T

	states []IFSMState[T]
	state  IFSMState[T]
}

func (fsm *FSM[T]) AddState(state IFSMState[T]) {
	fsm.states = append(fsm.states, state)

	hasState := fsm.state != nil

	if !hasState {
		fsm.state = state
		fsm.state.OnEnter(fsm.Target)
	}
}

func (fsm FSM[T]) GetState() IFSMState[T] {
	return fsm.state
}

func (fsm FSM[T]) getStateById(stateId FSMStateId) IFSMState[T] {
	for _, state := range fsm.states {
		isSearchingState := state.GetId() == stateId

		if isSearchingState {
			return state
		}
	}

	return nil
}

func (fsm *FSM[T]) SetState(stateId FSMStateId) bool {
	hasState := fsm.state != nil

	if !hasState {
		return false
	}

	newState := fsm.getStateById(stateId)
	hasNewState := newState != nil

	if !hasNewState {
		fmt.Printf("State %d is not part of FSM\n", stateId)
		return false
	}

	canChangeState := fsm.state.CanMoveTo(stateId)
	if !canChangeState {
		return false
	}

	fsm.state.OnExit(fsm.Target)
	newState.OnEnter(fsm.Target)
	fsm.state = newState

	return true
}

func (fsm *FSM[T]) Update() {
	fsm.state.Update(fsm.Target)
}
