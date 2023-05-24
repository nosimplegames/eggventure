package engine

type LivingCharacter struct {
	Character

	maxLife int
	life    int

	statusBar ICharacterStatusBar
}

func (character *LivingCharacter) SetStatusBar(statusBar ICharacterStatusBar) {
	character.statusBar = statusBar

	hasStatusBar := statusBar != nil

	if !hasStatusBar {
		return
	}

	statusBar.SetMaxHealth(character.maxLife)
	statusBar.SetHealth(character.life)
}

func (character LivingCharacter) HasStatusBar() bool {
	return character.statusBar != nil
}
func (character LivingCharacter) GetStatusBar() (ICharacterStatusBar, bool) {
	return character.statusBar, character.HasStatusBar()
}

type LivingCharacterFactory struct {
	CharacterFactory

	MaxLife int
	Life    int
}

func (factory LivingCharacterFactory) Init(character *LivingCharacter) {
	factory.CharacterFactory.Init(&character.Character)

	character.maxLife = factory.MaxLife
	character.life = factory.Life
}
