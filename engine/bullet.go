package engine

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/events"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/physics"
	"github.com/nosimplegames/ns-framework/render"
)

type Bullet struct {
	entities.Sprite
	MovementVector math.Vector

	ExplosionAnimation core.IAnimation
	isExploding        bool
}

func (bullet *Bullet) Update() {
	if bullet.isExploding {
		return
	}

	bullet.Move(bullet.MovementVector)
}

func (bullet Bullet) GetCollisionMask() string {
	return "bullet"
}

func (bullet Bullet) CanCollide() bool {
	return !bullet.isExploding
}

func (bullet Bullet) CanCollideWith(collisionMask string) bool {
	return collisionMask == "wall"
}

func (bullet *Bullet) OnCollision(collision physics.Collision) {
	xResolution := collision.CollisionResolverCalculator.CalculateXResolution()
	bullet.Move(math.Vector{
		X: xResolution,
	})
	bullet.explode()
}

func (bullet *Bullet) explode() {
	bullet.isExploding = true
	animation := bullet.ExplosionAnimation.Copy(bullet)
	core.AddAnimation(animation)
	animation.AddEventListener(events.EventListener{
		EventType: "stopped",
		Callback: func(_ events.Event) {
			bullet.Die()
		},
	})
}

type BulletFactory struct {
	Texture            render.Texture
	Position           math.Vector
	MovementVector     math.Vector
	ExplosionAnimation core.IAnimation
}

func (factory BulletFactory) Generate() {
	bullet := &Bullet{}
	entities.SpriteFactory{
		Texture: factory.Texture,
	}.Init(&bullet.Sprite)

	bullet.SetPosition(factory.Position)
	bullet.MovementVector = factory.MovementVector
	bullet.ExplosionAnimation = factory.ExplosionAnimation

	core.AddChildToRoot(bullet)
	physics.AddCollisionable(bullet)

	isGoingLeft := factory.MovementVector.X < 0

	if isGoingLeft {
		bullet.SetScale(math.Vector{
			X: -1,
			Y: 1,
		})
	}
}
