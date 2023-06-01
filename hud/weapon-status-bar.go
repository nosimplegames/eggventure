package hud

import (
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/render"
)

type WeaponStatusBar struct {
	core.Entity
	ammoSprite *entities.Sprite
}

func (bar WeaponStatusBar) SetAmmoTexture(texture render.Texture) {
	bar.ammoSprite.SetTexture(texture)
}

type WeaponStatusBarFactory struct {
	Ammo        int
	AmmoTexture render.Texture
}

func (factory WeaponStatusBarFactory) Create() *WeaponStatusBar {
	bar := &WeaponStatusBar{}

	ammoSprite := entities.SpriteFactory{
		Texture: factory.AmmoTexture,
	}.Create()
	ammoSprite.SetSize(res.AmmoTextureSize)
	ammoSprite.SetOrigin(res.AmmoTextureSize.By(0.5))
	bar.ammoSprite = ammoSprite

	core.EntityAdder{
		Parent: bar,
		Child:  ammoSprite,
	}.Add()

	return bar
}
