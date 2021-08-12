package main

import "fmt"

type Calendar struct {
	name string
}
type cont struct {
	a       int
	datamap map[string]Calendar
}

func newcont() *cont {
	calendar := make(map[string]Calendar)
	ctx := cont{
		a:       5,
		datamap: calendar,
	}
	return &ctx
}
func (c *cont) check() {
	c.a = 11
	c.datamap["r"] = Calendar{name: "dear"}
}

func (c *cont) prnt() {

	c.datamap["r"] = Calendar{name: "deart"}
	c.a++
	fmt.Println("Это")
}

func main() {

	mycntx := newcont()
	fmt.Println(mycntx)
	mycntx.check()
	fmt.Println(mycntx)
	mycntx.prnt()
	fmt.Println(mycntx)

}
