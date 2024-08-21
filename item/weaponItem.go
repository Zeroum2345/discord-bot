package item

type WeaponInterface interface {
	Item

	NewWeaponItem()
	BattleBehavior()
	ExplorationBehavior()
	IdleBehavior()
}

type WeaponItem struct {
	WeaponInterface

	Name       string
	Damage     float64
	Status     string
	LootRarity string
	DropChance float64
}

func (weaponItem WeaponItem) GetDropChance() float64 {
	return weaponItem.DropChance
}
