package main

// Board constructor
type Board struct {
	NailsNeeded int
	NailsDriven int
}

// NailDriver abstruction
type NailDriver interface {
	DriveNail(nailSuplly *int, b *Board)
}

// NailPuller abstruction
type NailPuller interface {
	PullNail(naiSupply *int, b *Board)
}

// NailController inherits from both above two I's
type NailController interface {
	NailDriver
	NailPuller
}

func main() {
	println("hellow world")
}
