package engine

import "github.com/nosimplegames/ns-framework/core"

type IInteractable interface {
	core.IEntity

	Interact(IInteractiveCharacter)
}
