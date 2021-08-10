package main

type builder2 struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *builder2 {
	return &builder2{}
}

func (b *builder2) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *builder2) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *builder2) setNumFloor() {
	b.floor = 1
}

func (b *builder2) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
