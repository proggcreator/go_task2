package main

type builder1 struct {
	walls string
	roof  string
}

func newBuilder1() *builder1 {
	return &builder1{}
}

func (b *builder1) setWallsType() {
	b.walls = "Дерево"
}

func (b *builder1) setRoofType() {
	b.roof = "Дерево"
}

func (b *builder1) getHouse() house {
	return house{
		roof:  b.roof,
		walls: b.walls,
	}
}
