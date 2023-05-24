package hud

import (
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
	"github.com/nosimplegames/ns-framework/ui"
)

type WeaponStatusBar struct {
	ui.Container
	ammoSprite *entities.Sprite
}

func (bar WeaponStatusBar) SetAmmoTexture(texture render.Texture) {
	bar.ammoSprite.SetTexture(texture)
	bar.ammoSprite.SetOriginCenter()

	bar.RepositionElements()
}

type WeaponStatusBarFactory struct {
	Ammo        int
	AmmoTexture render.Texture
}

func (factory WeaponStatusBarFactory) Create() *WeaponStatusBar {
	bar := &WeaponStatusBar{}

	bar.Layout = &ui.FlexLayout{
		Gap: math.Vector{
			X: res.WeaponStatusBarGap,
		},
	}

	ammoSprite := entities.SpriteFactory{
		Texture: factory.AmmoTexture,
	}.Create()
	bar.ammoSprite = ammoSprite
	bar.AddChild(ammoSprite)

	return bar
}
