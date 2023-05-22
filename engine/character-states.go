package engine

type CharacterState = FSMStateId

const (
	Idle CharacterState = iota
	Walking
	Jumping
	Falling
	Landing
)
