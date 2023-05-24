package hud

import (
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
	"github.com/nosimplegames/ns-framework/ui"
)

type HealthBar struct {
	ui.Container

	heartTexture      render.Texture
	emptyHeartTexture render.Texture

	health    int
	maxHealth int
}

func (bar *HealthBar) SetHealth(health int) {
	bar.health = health
	bar.updateHearts()
}

func (bar *HealthBar) SetMaxHealth(maxHealth int) {
	bar.maxHealth = maxHealth
	bar.updateHearts()
}

func (bar *HealthBar) updateHearts() {
	bar.RemoveChildren()
	bar.createHearts(bar.health)
	bar.createEmtpyHearts(bar.maxHealth - bar.health)
}

func (bar *HealthBar) createHearts(health int) {
	for i := 0; i < health; i++ {
		filledHeart := entities.SpriteFactory{
			Texture: bar.heartTexture,
		}.Create()

		bar.AddChild(filledHeart)
	}
}

func (bar *HealthBar) createEmtpyHearts(emptyHearts int) {
	for i := 0; i < emptyHearts; i++ {
		filledHeart := entities.SpriteFactory{
			Texture: bar.emptyHeartTexture,
		}.Create()

		bar.AddChild(filledHeart)
	}
}

type HealthBarFactory struct {
	MaxHealth int
	Health    int
}

func (factory HealthBarFactory) Create() *HealthBar {
	bar := &HealthBar{}

	bar.Layout = &ui.FlexLayout{
		Gap: math.Vector{X: res.HealthBarGap},
	}

	textures := res.GetTextures()
	bar.heartTexture = textures.FilledHeart
	bar.emptyHeartTexture = textures.EmptyHeart
	bar.health = factory.Health
	bar.maxHealth = factory.MaxHealth
	bar.updateHearts()

	bar.UseContentAsSize()

	return bar
}
