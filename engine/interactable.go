package engine

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/physics"
	"github.com/nosimplegames/ns-framework/render"
)

type Interactable struct {
	entities.Sprite

	interactionEntity core.IEntity
}

func (item Interactable) Update() {
	item.interactionEntity.Hide()
}

func (item Interactable) GetCollisionMask() string {
	return "interaction-item"
}

func (item Interactable) CanCollide() bool {
	return true
}

func (item Interactable) CanCollideWith(collisionMask string) bool {
	return collisionMask == "player"
}

func (item *Interactable) OnCollision(collision physics.Collision) {
	item.interactionEntity.Show()
}

type InteractableFactory struct {
	ItemTexture                render.Texture
	InteractionButtonAnimation core.IAnimation
	InteractionButtonSize      math.Vector
	InteractionButtonSpace     float64
}

func (factory InteractableFactory) Init(item *Interactable) {
	factory.initItem(item)

	interactionEntity := factory.createInteractionEntity()
	item.AddChild(interactionEntity)
	item.interactionEntity = interactionEntity
}

func (factory InteractableFactory) initItem(item *Interactable) {
	itemTextureSize := render.GetTextureSize(factory.ItemTexture)
	interactionArea := math.Vector{
		X: itemTextureSize.X * 2.0,
		Y: itemTextureSize.Y,
	}

	item.SetTexture(factory.ItemTexture)
	item.Size = interactionArea
	item.SetOriginCenter()
}

func (factory InteractableFactory) createInteractionEntity() core.IEntity {
	entity := &entities.Sprite{}
	itemSize := render.GetTextureSize(factory.ItemTexture)
	entity.SetPosition(math.Vector{
		X: itemSize.X * 0.5,
		Y: -factory.InteractionButtonSpace,
	})
	entity.SetOrigin(factory.InteractionButtonSize.By(0.5))
	entity.Hide()

	animation := factory.InteractionButtonAnimation.Copy(entity)
	core.AddAnimation(animation)

	return entity
}
