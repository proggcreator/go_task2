package main

type picturePng struct {
}

func (p *picturePng) accept(v visitor) {
	v.visitForPng(p)
}

func (c *picturePng) getType() string {
	return "PNG"
}
