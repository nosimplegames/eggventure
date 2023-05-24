package engine

import (
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
	"github.com/nosimplegames/ns-framework/utils"
)

type Weapon struct {
	entities.Sprite

	Ammo        int
	AmmoTexture render.Texture

	timeBetweenShoots  float64
	timeSinceLastShoot float64

	flash         *entities.Sprite
	flashDuration float64
}

func (weapon *Weapon) Update() {
	weapon.timeSinceLastShoot += utils.FrameTime

	needToHideFlash := weapon.flash.IsVisible() && weapon.timeSinceLastShoot >= weapon.flashDuration
	if needToHideFlash {
		weapon.flash.Hide()
	}
}

func (weapon *Weapon) Shoot() {
	canShoot := weapon.timeSinceLastShoot >= weapon.timeBetweenShoots

	if !canShoot {
		return
	}

	weapon.flash.Show()
	weapon.timeSinceLastShoot = 0
}

type WeaponFactory struct {
	Ammo        int
	AmmoTexture render.Texture

	Texture          render.Texture
	FlashTexture     render.Texture
	FlashAnchorPoint math.Vector
}

func (factory WeaponFactory) Create() *Weapon {
	weapon := &Weapon{}

	entities.SpriteFactory{
		Texture: factory.Texture,
	}.Init(&weapon.Sprite)

	weapon.Ammo = factory.Ammo
	weapon.AmmoTexture = factory.AmmoTexture

	flash := entities.SpriteFactory{
		Texture: factory.FlashTexture,
	}.Create()
	flash.Hide()
	flash.SetPosition(factory.FlashAnchorPoint)
	weapon.AddChild(flash)
	weapon.flash = flash
	weapon.flashDuration = 0.25

	weapon.timeBetweenShoots = 0.25

	return weapon
}
