package main

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
}

func main() {

	doctor := &doctor{}
	doctor.setNext(nil)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "John Smith"}
	reception.execute(patient)
	/*Reception registering patient
	Doctor checking patient*/
}
