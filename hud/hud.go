package hud

import (
	"github.com/nosimplegames/eggventure/engine"
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/ui"
)

type HUD struct {
	ui.Container

	statusBar engine.ICharacterStatusBar
}

func (hud HUD) GetStatusBar() engine.ICharacterStatusBar {
	return hud.statusBar
}

type HUDFactory struct {
	Size math.Vector
}

func (factory HUDFactory) Create() *HUD {
	hud := &HUD{}
	hud.Padding = res.HUDPadding
	hud.Layout = &ui.FlexLayout{}
	hud.Size = factory.Size

	playerStatusBar := PlayerStatusBarFactory{}.Create()
	hud.AddChild(playerStatusBar)
	hud.statusBar = playerStatusBar

	return hud
}
