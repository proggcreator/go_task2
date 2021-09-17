package main

import "fmt"

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	//if doctor checkup, execute next (now it's nil)
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true

}

func (d *doctor) setNext(next department) {
	d.next = next
}
