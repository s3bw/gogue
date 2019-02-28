package entity

type ItemSlots interface {
	SetItem(slot string, item Item)
	EmptySlot(slot string)
}

// map[string]Item
