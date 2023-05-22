package engine

type FSMStateId int

type IFSMState[T any] interface {
	GetId() FSMStateId
	OnEnter(target T)
	OnExit(target T)
	Update(target T)
	CanMoveTo(FSMStateId) bool
}
