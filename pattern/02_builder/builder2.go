package main

type builder2 struct {
	walls string
	roof  string
}

func newBuilder2() *builder1 {
	return &builder1{}
}

func (b *builder2) setWallsType() {
	b.walls = "Бетон"
}

func (b *builder2) setRoofType() {
	b.roof = "Стекло"
}

func (b *builder2) getHouse() house {
	return house{
		walls: b.walls,
		roof:  b.roof,
	}
}
