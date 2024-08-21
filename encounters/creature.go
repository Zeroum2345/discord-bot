package encounters

import "main/player"

type CreatureInterface interface {
	EncounterInterface

	NewCreature()
	TakeDamage()
	Attack()
	Chase()
}

type Creature struct {
	CreatureInterface
	Encounter

	Type   string
	Health float64
	Damage float64
}

func (creature Creature) TakeDamage(damage float64) {
	creature.Health -= damage
}

func (creature Creature) Attack(player player.Player) {
	player.Health -= creature.Damage
}
