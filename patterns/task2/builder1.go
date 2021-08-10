package main

type builder1 struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *builder1 {
	return &builder1{}
}

func (b *builder1) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *builder1) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *builder1) setNumFloor() {
	b.floor = 2
}

func (b *builder1) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
