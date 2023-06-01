package hud

import (
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/render"
)

type PlayerStatusBar struct {
	entities.Sprite

	healthBar       *HealthBar
	weaponStatusBar *WeaponStatusBar
}

func (bar *PlayerStatusBar) SetAmmo(ammo int) {
}

func (bar *PlayerStatusBar) SetAmmoTexture(texture render.Texture) {
	bar.weaponStatusBar.SetAmmoTexture(texture)
}

func (bar *PlayerStatusBar) SetHealth(health int) {
	bar.healthBar.SetHealth(health)
}

func (bar *PlayerStatusBar) SetMaxHealth(maxHealth int) {
	bar.healthBar.SetMaxHealth(maxHealth)
}

type PlayerStatusBarFactory struct {
}

func (factory PlayerStatusBarFactory) Create() *PlayerStatusBar {
	bar := &PlayerStatusBar{}

	entities.SpriteFactory{
		Texture: res.GetTextures().PlayerBar,
	}.Init(&bar.Sprite)

	characterIcon := factory.createCharacterIcon()

	healthBar := factory.createHealthBar()
	bar.healthBar = healthBar

	weaponStatusBar := factory.createWeaponBar()
	bar.weaponStatusBar = weaponStatusBar

	core.EntityAdder{
		Parent: bar,
		Children: core.EntityChildren{
			characterIcon,
			healthBar,
			weaponStatusBar,
		},
	}.Add()

	return bar
}

func (factory PlayerStatusBarFactory) createHealthBar() *HealthBar {
	bar := HealthBarFactory{
		MaxHealth: 3,
		Health:    3,
	}.Create()
	bar.SetPosition(res.HealthBarPosition)

	return bar
}

func (factory PlayerStatusBarFactory) createCharacterIcon() core.IEntity {
	icon := entities.SpriteFactory{
		Texture: res.GetTextures().EggCharacter,
	}.Create()
	icon.SetPosition(res.CharacterIconPosition)

	return icon
}

func (factory PlayerStatusBarFactory) createWeaponBar() *WeaponStatusBar {
	bar := WeaponStatusBarFactory{}.Create()
	bar.SetPosition(res.WeaponStatusBarPosition)

	return bar
}

// func (factory PlayerStatusBarFactory) createContent(bar *PlayerStatusBar) core.IEntity {
// 	content := &ui.Container{}

// 	size := bar.GetSize()
// 	contentGap := 5.0
// 	content.SetSize(size)
// 	content.Padding = res.PlayerBarPadding
// 	content.SetPosition(size.By(0.5))
// 	content.Layout = &ui.FlexLayout{
// 		Gap: math.Vector{X: contentGap},
// 	}

// 	character := entities.SpriteFactory{
// 		Texture: res.GetTextures().EggCharacter,
// 	}.Create()
// 	content.AddChild(character)

// 	leftContent := factory.getLeftContent(bar)
// 	content.AddChild(leftContent)

// 	return content
// }

// func (factory PlayerStatusBarFactory) getLeftContent(bar *PlayerStatusBar) core.IEntity {
// 	leftContent := &ui.Container{}
// 	leftContent.Layout = &ui.FlexLayout{
// 		LayoutDirection: ui.FlexColumn,
// 		Gap:             math.Vector{X: res.SpaceBetweenHealthAndWeaponBars},
// 	}

// 	healthBar := HealthBarFactory{
// 		MaxHealth: 3,
// 		Health:    3,
// 	}.Create()
// 	leftContent.AddChild(healthBar)
// 	bar.healthBar = healthBar

// 	weaponStatusBar := WeaponStatusBarFactory{}.Create()
// 	bar.AddChild(weaponStatusBar)
// 	bar.weaponStatusBar = weaponStatusBar

// 	return leftContent
// }
