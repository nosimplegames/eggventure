package engine

import (
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type DamageMask struct {
	entities.Sprite

	lowDamageTexture    render.Texture
	mediumDamageTexture render.Texture
	highDamageTexture   render.Texture

	maxHealth int
	health    int
}

func (mask *DamageMask) SetHealth(health int) {
	mask.health = health

	mask.updateTexture()
}

func (mask *DamageMask) updateTexture() {
	healthPercentage := float64(mask.health) / float64(mask.maxHealth)

	if healthPercentage >= 0.9 {
		return
	} else if healthPercentage >= 0.7 {
		mask.SetTexture(mask.lowDamageTexture)
	} else if healthPercentage >= 0.4 {
		mask.SetTexture(mask.mediumDamageTexture)
	} else {
		mask.SetTexture(mask.highDamageTexture)
	}
}

type DamageMaskFactory struct {
	Size math.Vector

	LowDamageTexture    render.Texture
	MediumDamageTexture render.Texture
	HighDamageTexture   render.Texture

	MaxHealth int
	Health    int
}

func (factory DamageMaskFactory) Create() *DamageMask {
	mask := &DamageMask{}

	mask.SetSize(factory.Size)
	mask.SetOrigin(factory.Size.By(0.5))
	mask.lowDamageTexture = factory.LowDamageTexture
	mask.mediumDamageTexture = factory.MediumDamageTexture
	mask.highDamageTexture = factory.HighDamageTexture

	mask.maxHealth = factory.MaxHealth
	mask.health = factory.Health
	mask.updateTexture()

	return mask
}
