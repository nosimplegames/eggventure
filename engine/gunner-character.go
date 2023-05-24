package engine

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
)

type GunnerCharacter struct {
	LivingCharacter

	weapon            *Weapon
	weaponAnchorPoint math.Vector
}

func (character *GunnerCharacter) PickUpWeapon(weapon *Weapon) {
	character.weapon = weapon
	weapon.SetPosition(character.weaponAnchorPoint)
	character.AddChild(weapon)

	character.reportWeaponInfo()
}

func (character GunnerCharacter) Shoot() {
	hasWeapon := character.weapon != nil

	if !hasWeapon {
		return
	}

	character.weapon.Shoot()
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

type GunnerCharacterFactory struct {
	LivingCharacterFactory

	WeaponAnchorPoint math.Vector
}

func (factory GunnerCharacterFactory) Init(character *GunnerCharacter) {
	factory.LivingCharacterFactory.Init(&character.LivingCharacter)
	character.weaponAnchorPoint = factory.WeaponAnchorPoint
	character.SetDrawPolicy(core.DrawAfterChildren)
}
