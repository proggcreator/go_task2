package main

type iBuilder interface {
	setWallsType()
	setRoofType()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "house1" {
		return &builder1{}
	}

	if builderType == "house2" {
		return &builder2{}
	}
	return nil
}
