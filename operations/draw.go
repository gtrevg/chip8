package operations

import (
	"fmt"
	"github.com/jtharris/chip8/system"
	"math/bits"
)

// Parser for DrawOp
type drawParser struct{}

func (p drawParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xD
}

func (p drawParser) createOp(opcode system.OpCode) Operation {
	return DrawOp{
		xRegister: byte(opcode & 0x0F00 >> 8),
		yRegister: byte(opcode & 0x00F0 >> 4),
		height:    byte(opcode & 0x000F),
	}
}

// DrawOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Dxyn
type DrawOp struct {
	xRegister byte
	yRegister byte
	height    byte
}

// String returns a text representation of this operation
func (o DrawOp) String() string {
	return fmt.Sprintf("Draw Screen (V%X, V%X) Height: %X", o.xRegister, o.yRegister, o.height)
}

// Execute this operation on the given virtual machine
func (o DrawOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[0xF] = 0 // start with this as the default position
	xPos := vm.Registers[o.xRegister]
	yPos := vm.Registers[o.yRegister]

	for row := byte(0); row < o.height; row++ {
		y := (yPos + row) % 32

		sprite := uint64(vm.Memory[vm.IndexRegister+uint16(row)])
		sprite = bits.RotateLeft64(sprite, 56-int(xPos))

		// If any 'on' pixels are going to be flipped, then set
		// VF to 1 per the spec
		if sprite&vm.Pixels[y] > 0 {
			vm.Registers[0xF] = 1
		}

		vm.Pixels[y] ^= sprite
	}
}
