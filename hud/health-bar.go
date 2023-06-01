package hud

import (
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
)

type HealthBar struct {
	core.Entity

	health    int
	maxHealth int
	hearts    []*Heart
}

func (bar *HealthBar) SetHealth(health int) {
	bar.health = health

	needToCapHealth := bar.health > bar.maxHealth

	if needToCapHealth {
		bar.health = bar.maxHealth
	}

	bar.emptyHearts()
	bar.fillHearts(bar.health)
}

func (bar *HealthBar) SetMaxHealth(maxHealth int) {
	needToCreateHearts := bar.maxHealth < maxHealth

	if needToCreateHearts {
		bar.createHearts(maxHealth - bar.maxHealth)
		bar.maxHealth = maxHealth
	}
}

func (bar HealthBar) emptyHearts() {
	for _, heart := range bar.hearts {
		heart.Empty()
	}
}

func (bar HealthBar) fillHearts(count int) {
	for i := 0; i < count; i++ {
		heart := bar.hearts[i]
		heart.Fill()
	}
}

// func (bar *HealthBar) updateHearts() {
// 	bar.RemoveChildren()
// 	bar.createHearts(bar.health)
// 	bar.createEmtpyHearts(bar.maxHealth - bar.health)
// }

func (bar *HealthBar) createHearts(count int) {
	heartWidth := (res.HeartSize.X + res.HeartMargin)
	heartPosition := math.Vector{
		X: float64(len(bar.hearts)) * heartWidth,
	}

	for i := 0; i < count; i++ {
		heart := HeartFactory{
			IsHeartEmpty: true,
		}.Create()

		heart.SetPosition(heartPosition)
		heartPosition.X += heartWidth

		bar.hearts = append(bar.hearts, heart)
		core.EntityAdder{
			Parent: bar,
			Child:  heart,
		}.Add()
	}
}

// func (bar *HealthBar) createEmtpyHearts(emptyHearts int) {
// 	for i := 0; i < emptyHearts; i++ {
// 		filledHeart := entities.SpriteFactory{
// 			Texture: bar.emptyHeartTexture,
// 		}.Create()

// 		bar.AddChild(filledHeart)
// 	}
// }

type HealthBarFactory struct {
	MaxHealth int
	Health    int
}

func (factory HealthBarFactory) Create() *HealthBar {
	bar := &HealthBar{}
	bar.SetMaxHealth(factory.MaxHealth)
	bar.SetHealth(factory.Health)

	// bar.Layout = &ui.FlexLayout{
	// 	Gap: math.Vector{X: res.HealthBarGap},
	// }

	// bar.health = factory.Health
	// bar.maxHealth = factory.MaxHealth
	// bar.updateHearts()

	return bar
}
