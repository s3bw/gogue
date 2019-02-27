package gogue

import (
	"fmt"

	"github.com/foxyblue/gogue/gogue/entity"
)

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

func (a *Area) GetItems() []*entity.Item {
	var list []*entity.Item

	for _, i := range a.Entities {
		if i.Identify() == entity.TypeItem {
			list = append(list, i.(*entity.Item))
		}
	}
	return list
}

func RemoveFromMap(target *entity.Item, entities []entity.Entity) []entity.Entity {
	var targetIndex int

	for index, m := range entities {
		if target == m {
			targetIndex = index
			break
		}
	}
	return append(entities[:targetIndex], entities[targetIndex+1:]...)
}

func (a *Area) Pickup(obj *entity.Creature) bool {
	var target *entity.Item

	x := obj.X
	y := obj.Y
	target = nil

	for _, item := range a.GetItems() {
		if item.X == x && item.Y == y {
			target = item
			break
		}
	}
	if target == nil {
		return false
	} else if obj.CanCarry(target.Weight) {
		a.Entities = RemoveFromMap(target, a.Entities)
		a.Feed.Log(fmt.Sprintf("Picked up %s!", target.Name))
		return true
	}
	a.Feed.Log(fmt.Sprintf("%s is too heavy!", target.Name))
	return false
}

// MoveCreature will move a creature in the area, checking for collisions
func (a *Area) MoveCreature(obj *entity.Creature, dx, dy int) {
	var target *entity.Creature

	x := obj.X + dx
	y := obj.Y + dy
	target = nil
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
		obj.Attack(target, a.Feed)
	}
}
