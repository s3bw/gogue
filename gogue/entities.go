package gogue

import "github.com/foxyblue/gogue/gogue/entity"

// GetCreatures is effectively an iterator
func (a *Area) GetCreatures() []*entity.Creature {
	var list []*entity.Creature

	for _, i := range a.Entities {
		if i.Identify() == entity.TypeCreature {
			list = append(list, i.(*entity.Creature))
		}
	}
	return list
}

// MoveCreature will move a creature in the area, checking for collisions
func (a *Area) MoveCreature(obj *entity.Creature, dx, dy int) {
	var target *entity.Creature

	x := obj.X + dx
	y := obj.Y + dy
	target = nil
	// I'm going to have to use an iterator which filters
	// for types.
	for _, monster := range a.GetCreatures() {
		if monster.X == x && monster.Y == y {
			target = monster
			break
		}
	}
	if target == nil {
		if a.Grid.Tiles[x][y].Passable {
			obj.Move(x, y)
		}
	} else {
		// This can kill, when a creature dies, it transforms
		// into an object. I won't have to remove it from the
		// list. I might have to replace it.
		obj.Attack(target, a.Feed)
	}
}
