package engine

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/physics"
	"github.com/nosimplegames/ns-framework/render"
)

type LivingCharacter struct {
	Character

	maxHealth int
	health    int

	statusBar ICharacterStatusBar

	damageMask *DamageMask
}

func (character *LivingCharacter) SetStatusBar(statusBar ICharacterStatusBar) {
	character.statusBar = statusBar

	hasStatusBar := statusBar != nil

	if !hasStatusBar {
		return
	}

	statusBar.SetMaxHealth(character.maxHealth)
	statusBar.SetHealth(character.health)
}

func (character LivingCharacter) HasStatusBar() bool {
	return character.statusBar != nil
}

func (character LivingCharacter) GetStatusBar() (ICharacterStatusBar, bool) {
	return character.statusBar, character.HasStatusBar()
}

func (character LivingCharacter) CanCollideWith(collisionMask string) bool {
	return character.Character.CanCollideWith(collisionMask) ||
		collisionMask == "bullet"
}

func (character *LivingCharacter) OnCollision(collision physics.Collision) {
	character.Character.OnCollision(collision)

	switch collision.AnotherCollisionMask {
	case "bullet":
		character.OnBulletCollision(collision)
	}
}

func (character *LivingCharacter) OnBulletCollision(collision physics.Collision) {
	character.SetHealth(character.health - 1)
}

func (character *LivingCharacter) SetHealth(health int) {
	character.health = health
	character.damageMask.SetHealth(health)
	statusBar, hasStatusBar := character.GetStatusBar()

	if !hasStatusBar {
		return
	}

	statusBar.SetHealth(health)
}

type LivingCharacterFactory struct {
	CharacterFactory

	Size math.Vector

	MaxHealth int
	Health    int

	DamageMaskSize      math.Vector
	LowDamageTexture    render.Texture
	MediumDamageTexture render.Texture
	HighDamageTexture   render.Texture
}

func (factory LivingCharacterFactory) Init(character *LivingCharacter) {
	factory.CharacterFactory.Init(&character.Character)

	character.maxHealth = factory.MaxHealth
	character.health = factory.Health

	damageMask := factory.createDamageMask()
	core.EntityAdder{
		Parent: character,
		Child:  damageMask,
	}.Add()
	character.damageMask = damageMask
}

func (factory LivingCharacterFactory) createDamageMask() *DamageMask {
	mask := DamageMaskFactory{
		Size: factory.DamageMaskSize,

		LowDamageTexture:    factory.LowDamageTexture,
		MediumDamageTexture: factory.MediumDamageTexture,
		HighDamageTexture:   factory.HighDamageTexture,

		MaxHealth: factory.MaxHealth,
		Health:    factory.Health,
	}.Create()

	mask.SetPosition(factory.Size.By(0.5))

	return mask
}
