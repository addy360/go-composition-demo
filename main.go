package main

import "fmt"

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

// Mallet hammerlike
type Mallet struct {
}

// DriveNail drilling nails to a specified board
func (Mallet) DriveNail(nailSupply *int, b *Board) {

	*nailSupply--
	b.NailsDriven++

	fmt.Printf("Mallet drilled %d nails into a board, Supply remained with %d", b.NailsDriven, *nailSupply)
}

// Crowber for removing nails
type Crowber struct {
}

// NailPuller for removing nails from board
func (Crowber) NailPuller(nailSupply *int, b *Board) {
	*nailSupply++
	b.NailsDriven--
	fmt.Printf("Crowber removed a nail, %d nails remain into a board, Supply remained with %d", b.NailsDriven, *nailSupply)
}

// Contractor like a person using the two defined tools to secure boards
type Contractor struct {
}

// Fasten the actual driving nails into a board
func (Contractor) Fasten(d NailDriver, nailSuplly *int, b *Board) {
	for b.NailsDriven < b.NailsNeeded {
		d.DriveNail(nailSuplly, b)
	}
}

// Unfasten removes nailes from a board
func (Contractor) Unfasten(p NailPuller, nailSupply *int, b *Board) {
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}

func main() {
	println("hellow world")
}
