package player

import "main/item"

type Player struct {
	ID        string
	Name      string
	Level     int
	MaxHealth float64
	Health    float64
	Inventory []item.Item
	Equipped  []item.WeaponItem
}
