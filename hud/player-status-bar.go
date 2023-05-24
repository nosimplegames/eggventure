package hud

import (
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
	"github.com/nosimplegames/ns-framework/ui"
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

	bar.SetTexture(res.GetTextures().PlayerBar)
	bar.UseTextureSizeAsSize()
	bar.SetOriginCenter()

	content := factory.createContent(bar)
	bar.AddChild(content)

	return bar
}

func (factory PlayerStatusBarFactory) createContent(bar *PlayerStatusBar) core.IEntity {
	content := &ui.Container{}

	contentGap := 5.0
	content.Size = bar.Size
	content.Padding = res.PlayerBarPadding
	content.SetPosition(bar.Size.By(0.5))
	content.Layout = &ui.FlexLayout{
		Gap: math.Vector{X: contentGap},
	}

	character := &entities.Sprite{}
	character.Texture = res.GetTextures().EggCharacter
	character.UseTextureSizeAsSize()
	character.SetOriginCenter()
	content.AddChild(character)

	leftContent := factory.getLeftContent(bar)
	content.AddChild(leftContent)

	return content
}

func (factory PlayerStatusBarFactory) getLeftContent(bar *PlayerStatusBar) core.IEntity {
	leftContent := &ui.Container{}
	leftContent.Layout = &ui.FlexLayout{
		LayoutDirection: ui.FlexColumn,
		Gap:             math.Vector{X: res.SpaceBetweenHealthAndWeaponBars},
	}

	healthBar := HealthBarFactory{
		MaxHealth: 3,
		Health:    3,
	}.Create()
	leftContent.AddChild(healthBar)
	bar.healthBar = healthBar

	weaponStatusBar := WeaponStatusBarFactory{}.Create()
	bar.AddChild(weaponStatusBar)
	bar.weaponStatusBar = weaponStatusBar

	leftContent.UseContentAsSize()

	return leftContent
}
