package engine

type IFSM[T any] interface {
	AddState(IFSMState[T])
	SetState(stateId FSMStateId) bool
	Update()
}
