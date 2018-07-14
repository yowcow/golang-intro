package mybinary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinShift(t *testing.T) {
	a := 1 << 3
	b := 8 >> 1

	assert.Equal(t, 8, a)
	assert.Equal(t, 4, b)
}

func TestBinANDOR(t *testing.T) {
	a := 1
	b := 2
	and := a & b
	or := a | b

	assert.Equal(t, 0, and)
	assert.Equal(t, 3, or)
}

func TestOnesComplement(t *testing.T) {
	a := 8
	b := ^a

	assert.Equal(t, -9, b)
}

/*
	12   => 0000 1100
	12-1 => 0000 1011
	-----------------
	&       0000 1000 (8)
	8-1  => 0000 0111
	-----------------
	&    => 0000 0000 (0)
	That's it!
*/
func countBits(x uint8) (count int) {
	for count = 0; x != 0; x = x & (x - 1) {
		count++
	}
	return
}

func TestCountBits(t *testing.T) {
	// 0011 => 2
	assert.Equal(t, 2, countBits(0x3))
	// 0100 => 1
	assert.Equal(t, 1, countBits(0x4))
	// 1100 => 1
	assert.Equal(t, 2, countBits(0xC))
	// 1111 => 4
	assert.Equal(t, 4, countBits(0xF))
}

func TestComplement(t *testing.T) {
	x := uint8(3)
	y := ^x

	assert.Equal(t, "11111100", fmt.Sprintf("%b", y))
}

/*
	To invert `1100`, from 1 with padding 3,
	all I need is `1110` and XOR.

	To make `1110`,
	1. make 1111 (^0)
	2. shift right with padding width to make 1000
	3. invert 1000 and shift with starting index to make 1110
*/
func invertBits(x uint8, from, pad uint) (res uint8) {
	var mask uint8
	mask = ^(^uint8(0) << pad) << from
	res = x ^ mask
	return
}

func TestInvertBits(t *testing.T) {
	// given:  1100, invert bits from 1 with padding 3
	// result: 0010
	assert.Equal(t, uint8(2), invertBits(0xC, 1, 3))

	// given:  0001, invert bits from 0 with padding 2
	// result: 0010
	assert.Equal(t, uint8(2), invertBits(0x1, 0, 2))

	// given:  0001, invert bits from 0 with padding 3
	// result: 1110
	assert.Equal(t, uint8(14), invertBits(0x1, 0, 4))
}

func TestBigEndianHexToUint32Conversion(t *testing.T) {
	data := []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x03,
	}
	r := bytes.NewBuffer(data)

	for i := 0; i < 4; i++ {
		b := make([]byte, 4)
		n, err := r.Read(b)

		assert.Equal(t, 4, n)
		assert.Nil(t, err)

		v := binary.BigEndian.Uint32(b)

		assert.Equal(t, uint32(i), v)
	}
}

func TestLittleEndianHexToUint32Conversion(t *testing.T) {
	data := []byte{
		0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,
		0x02, 0x00, 0x00, 0x00,
		0x03, 0x00, 0x00, 0x00,
	}
	r := bytes.NewBuffer(data)

	for i := 0; i < 4; i++ {
		b := make([]byte, 4)
		n, err := r.Read(b)

		assert.Equal(t, 4, n)
		assert.Nil(t, err)

		v := binary.LittleEndian.Uint32(b)

		assert.Equal(t, uint32(i), v)
	}
}
