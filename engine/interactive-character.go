package engine

import "github.com/nosimplegames/ns-framework/physics"

type InteractiveCharacter struct {
	GunnerCharacter

	Interactable IInteractable
}

func (character *InteractiveCharacter) Update() {
	character.GunnerCharacter.Update()

	character.Interactable = nil
}

func (character InteractiveCharacter) CanCollideWith(collisionMask string) bool {
	return character.GunnerCharacter.CanCollideWith(collisionMask) ||
		collisionMask == "interaction-item"
}

func (character *InteractiveCharacter) OnCollision(collision physics.Collision) {
	switch collision.AnotherCollisionMask {
	case "interaction-item":
		character.Interactable = collision.Another.(IInteractable)
		return
	}

	character.GunnerCharacter.OnCollision(collision)
}

func (character *InteractiveCharacter) Interact() {
	hasInteractionItem := character.Interactable != nil

	if !hasInteractionItem {
		return
	}

	character.Interactable.Interact(character)
}

type InteractiveCharacterFactory struct {
	GunnerCharacterFactory
}

func (factory InteractiveCharacterFactory) Init(character *InteractiveCharacter) {
	factory.GunnerCharacterFactory.Init(&character.GunnerCharacter)
}
