package hud

import (
	"github.com/nosimplegames/eggventure/engine"
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
)

type HUD struct {
	core.Entity

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
	hud.SetSize(factory.Size)
	hud.SetOrigin(factory.Size.By(0.5))

	playerStatusBar := PlayerStatusBarFactory{}.Create()
	playerStatusBar.SetPosition(res.PlayerStatusBarPosition)
	core.EntityAdder{
		Parent: hud,
		Child:  playerStatusBar,
	}.Add()
	hud.statusBar = playerStatusBar

	return hud
}
