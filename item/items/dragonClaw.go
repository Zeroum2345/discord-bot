package items

import "main/item"

type DragonClawItem struct {
	item.WeaponItem
}

func NewWeaponItem() *DragonClawItem {
	return &DragonClawItem{
		item.WeaponItem{
			Name:       "Dragon Claw",
			Damage:     10.0,
			Status:     "Burn",
			LootRarity: "Epic",
			DropChance: 0.18,
		},
	}
}

var DragonClaw = NewWeaponItem()
