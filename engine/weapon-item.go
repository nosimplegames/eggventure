package engine

type WeaponItem struct {
	Interactable

	weaponFactory WeaponFactory
}

func (item *WeaponItem) Interact(character IInteractiveCharacter) {
	character.PickUpWeapon(item.createWeapon())
	item.Die()
}

func (item WeaponItem) createWeapon() *Weapon {
	weapon := item.weaponFactory.Create()

	return weapon
}

type WeaponItemFactory struct {
	InteractableFactory
	WeaponFactory
}

func (factory WeaponItemFactory) Create() *WeaponItem {
	item := &WeaponItem{}
	factory.InteractableFactory.Init(&item.Interactable)

	item.weaponFactory = factory.WeaponFactory

	return item
}
