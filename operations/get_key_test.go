package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestGetKeyParser_Matches(t *testing.T) {
	parser := GetKeyParser{}

	assert.True(t, parser.Matches(0xF30A))
}

func TestGetKeyParser_DoesNotMatch(t *testing.T) {
	parser := GetKeyParser{}

	assert.False(t, parser.Matches(0xF32A))
}

func TestGetKeyParser_CreateOp(t *testing.T) {
	parser := GetKeyParser{}
	expected := GetKeyOp{ register: 0xC }

	assert.Equal(t, expected, parser.CreateOp(0xFC0A))
}

func TestGetKeyOp_String(t *testing.T) {
	op := GetKeyOp{ register: 0x9 }

	assert.Equal(t, "V9 = get_key()", op.String())
}

func TestGetKeyOp_Execute(t *testing.T) {
	// Given
	vm := system.NewVirtualMachine()
	vm.Keyboard[0xA] = true
	op := GetKeyOp{ register: 0x6}

	// When
	op.Execute(&vm)

	// Then
	assert.True(t, vm.Running)
	assert.Equal(t, byte(0xA), vm.Registers[0x6])
}