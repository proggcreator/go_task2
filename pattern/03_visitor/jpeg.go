package main

type pictureJpeg struct {
}

func (p *pictureJpeg) accept(v visitor) {
	v.visitForJpeg(p)
}

func (p *pictureJpeg) getType() string {
	return "JPEG"
}
