package encounters

import (
	"main/item"
)

type EncounterInterface interface {
	Interact()
	GenerateLoot() item.Item
}

type Encounter struct {
	EncounterInterface
	LootTable item.LootTable
}
