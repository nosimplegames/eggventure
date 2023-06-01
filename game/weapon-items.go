package game

import (
	"github.com/nosimplegames/eggventure/engine"
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/math"
)

type WeaponItemsFactory struct {
}

func (factory WeaponItemsFactory) CreateMagnum() *engine.WeaponItem {
	textures := res.GetTextures()
	animations := res.GetAnimations()

	item := engine.WeaponItemFactory{
		WeaponFactory: engine.WeaponFactory{
			Ammo:         256,
			AmmoTexture:  textures.MagnumAmmo,
			Texture:      textures.Magnum,
			FlashTexture: textures.MagnumFlash,
			FlashAnchorPoint: math.Vector{
				X: 15,
				Y: 6,
			},
			AmmoExplosionAnimation: animations.MagnumExplosionAnimation,
		},

		InteractableFactory: engine.InteractableFactory{
			ItemTexture:                textures.MagnumItem,
			InteractionButtonAnimation: animations.ActionButtonAnimation,
			InteractionButtonSize:      res.InteractionButtonSize,
			InteractionButtonSpace:     res.InteractionButtonSpace,
		},
	}.Create()

	return item
}
