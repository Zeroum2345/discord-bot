package item

import "strconv"

type LootTable struct {
	Items []Item
	Loot  map[string]Item
}

func NewLootTable(items []Item) *LootTable {
	loot := make(map[string]Item)
	chance := 0.0
	totalProb := 0.0

	for _, item := range items {
		totalProb += item.GetDropChance()
	}

	for _, item := range items {
		chance += item.GetDropChance() / totalProb
		loot[strconv.FormatFloat(chance, 'f', -1, 64)] = item
	}

	return &LootTable{
		Items: items,
		Loot:  loot,
	}
}

func (lootTable LootTable) GetLootByChance(chance float64) Item {
	var chosenLoot Item

	for key, loot := range lootTable.Loot {
		lootChance, _ := strconv.ParseFloat(key, 64)
		if chance <= lootChance {
			chosenLoot = loot
			break
		}
	}

	return chosenLoot
}
