package main

import (
	"fmt"
	"log"
)

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

	fmt.Printf("Mallet drilled %d nails into a board, Supply increased to %d\n", b.NailsDriven, *nailSupply)
}

// Crowber for removing nails
type Crowber struct {
}

// PullNail for removing nails from board
func (Crowber) PullNail(nailSupply *int, b *Board) {
	*nailSupply++
	b.NailsDriven--
	fmt.Printf("Crowber removed a nail, %d nails remain into a board, Supply remained with %d\n", b.NailsDriven, *nailSupply)
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

// ProcessBoards allowing a contractor to process more than a board at a time
func (c Contractor) ProcessBoards(pd NailController, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("\nContractor checking with board #%d : %+v\n", i+1, b)

		switch {
		case b.NailsDriven < b.NailsNeeded:
			c.Fasten(pd, nailSupply, b)
		case b.NailsDriven > b.NailsNeeded:
			c.Unfasten(pd, nailSupply, b)
		}
	}
}

// ToolBox for a contractor
type ToolBox struct {
	NailDriver
	NailPuller
	nails int
}

func main() {

	boards := []Board{
		// Rotted boards to be removed.
		{NailsDriven: 5},
		{NailsDriven: 3},
		{NailsDriven: 1},

		// Fresh boards to be fasten.
		{NailsNeeded: 6},
		{NailsNeeded: 8},
		{NailsNeeded: 7},
	}

	tb := ToolBox{
		nails:      20,
		NailDriver: Mallet{},
		NailPuller: Crowber{},
	}
	log.Println("\nBefore processing")
	display(&tb, boards)

	var c Contractor
	c.ProcessBoards(&tb, &tb.nails, boards)

	log.Println("\nAfter processing")
	display(&tb, boards)
}

func display(tb *ToolBox, boards []Board) {
	fmt.Printf("ToolBox %#v\n", tb)
	log.Println("Boards")
	defer fmt.Println()
	for _, v := range boards {
		fmt.Printf("[*] %+v\n", v)
	}
}
