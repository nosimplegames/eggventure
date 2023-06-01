package engine

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
	"github.com/nosimplegames/ns-framework/utils"
)

type Weapon struct {
	entities.Sprite

	Ammo                   int
	AmmoTexture            render.Texture
	AmmoExplosionAnimation core.IAnimation

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

func (weapon *Weapon) Shoot(direction MovingDirection) {
	canShoot := weapon.timeSinceLastShoot >= weapon.timeBetweenShoots

	if !canShoot {
		return
	}

	weapon.flash.Show()
	weapon.timeSinceLastShoot = 0
	weapon.generateBullet(direction)
}

func (weapon Weapon) generateBullet(direction MovingDirection) {
	position := weapon.GetPosition()
	movementVector := math.MovementVector{
		Speed: 6.5,
	}.Calculate()

	isGoingLeft := direction == MovingDirectionLeft

	if isGoingLeft {
		movementVector.X = -movementVector.X
	}

	BulletFactory{
		Texture:            weapon.AmmoTexture,
		Position:           position,
		MovementVector:     movementVector,
		ExplosionAnimation: weapon.AmmoExplosionAnimation,
	}.Generate()
}

type WeaponFactory struct {
	Ammo                   int
	AmmoTexture            render.Texture
	AmmoExplosionAnimation core.IAnimation

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
	weapon.AmmoExplosionAnimation = factory.AmmoExplosionAnimation

	flash := entities.SpriteFactory{
		Texture: factory.FlashTexture,
	}.Create()
	flash.Hide()
	flash.SetPosition(factory.FlashAnchorPoint)
	core.EntityAdder{
		Parent: weapon,
		Child:  flash,
	}.Add()
	weapon.flash = flash
	weapon.flashDuration = 0.1

	weapon.timeBetweenShoots = 0.5
	weapon.timeSinceLastShoot = weapon.timeBetweenShoots

	return weapon
}
