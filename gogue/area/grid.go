package area

type Block struct {
	visible    bool
	passable   bool
	appearance rune
}

type Grid struct {
	titles [][]*Block
}
