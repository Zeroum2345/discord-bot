package creatures

import (
	"main/encounters"
	"main/item"
	"main/item/items"
)

type RedDragonCreature struct {
	encounters.Creature
	encounters.Encounter
}

func NewCreature() *RedDragonCreature {
	lootTable := *item.NewLootTable([]item.Item{items.DragonClaw})

	return &RedDragonCreature{
		Encounter: encounters.Encounter{
			LootTable: lootTable,
		},
		Creature: encounters.Creature{
			Type:   "Dragon",
			Health: 100.0,
			Damage: 23.0,
		},
	}
}

var RedDragon = NewCreature()
