package engine

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
)

type GunnerCharacter struct {
	LivingCharacter

	weapon            *Weapon
	weaponAnchorPoint math.Vector
	aimingDirection   MovingDirection
}

func (character *GunnerCharacter) PickUpWeapon(weapon *Weapon) {
	character.weapon = weapon
	weapon.SetPosition(character.weaponAnchorPoint)
	core.EntityAdder{
		Parent: character,
		Child:  weapon,
	}.Add()

	character.reportWeaponInfo()
}

func (character GunnerCharacter) Shoot() {
	hasWeapon := character.weapon != nil

	if !hasWeapon {
		return
	}

	character.weapon.Shoot(character.aimingDirection)
}

func (character *GunnerCharacter) SetStatusBar(statusBar ICharacterStatusBar) {
	character.LivingCharacter.SetStatusBar(statusBar)
	character.reportWeaponInfo()
}

func (character GunnerCharacter) reportWeaponInfo() {
	characterStatusBar, hasStatusBar := character.GetStatusBar()
	canSetWeaponInfo := hasStatusBar && character.weapon != nil

	if !canSetWeaponInfo {
		return
	}

	characterStatusBar.SetAmmoTexture(character.weapon.AmmoTexture)
	characterStatusBar.SetAmmo(character.weapon.Ammo)
}

func (character *GunnerCharacter) SetMovingDirection(direction MovingDirection) {
	character.LivingCharacter.SetMovingDirection(direction)

	isValidAimingDirection := direction != NoMoving

	if !isValidAimingDirection {
		return
	}

	character.aimingDirection = direction
}

type GunnerCharacterFactory struct {
	LivingCharacterFactory

	WeaponAnchorPoint math.Vector
}

func (factory GunnerCharacterFactory) Init(character *GunnerCharacter) {
	factory.LivingCharacterFactory.Init(&character.LivingCharacter)
	character.weaponAnchorPoint = factory.WeaponAnchorPoint
	character.SetDrawPolicy(core.DrawAfterChildren)
}
