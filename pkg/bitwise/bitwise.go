package bitwise

type Integer interface {
	int | int64 | int32 | byte | uint | uint64 | uint16
}

// GetMask return a bit mask : all bits below bitNum set to 1
func GetMask[K Integer](bitNum byte) K {
	return (1 << bitNum) - 1
}

// GetBit return True if bit is set or False in opposite case
func GetBit[K Integer](value K, bitNum byte) bool {
	if value&(1<<bitNum) > 0 {
		return true
	}
	return false
}

// SetBit sets bit to 1 in the provided value
func SetBit[K Integer](value K, bitNum byte) K {
	return value | (1 << bitNum)
}

// ClearBit sets bit to 0 in the provided value
func ClearBit[K Integer](value K, bitNum byte) K {
	return value & ^(1 << bitNum)
}

// BitCount returns number of set bits in the val
func BitCount(val int) int {
	val = ((val >> 1) & 0b01010101010101010101010101010101) +
		(val & 0b01010101010101010101010101010101)
	val = ((val >> 2) & 0b00110011001100110011001100110011) +
		(val & 0b00110011001100110011001100110011)
	val = (val>>4)&0b00001111000011110000111100001111 +
		(val & 0b00001111000011110000111100001111)
	val = (val>>8)&0b00000000111111110000000011111111 +
		(val & 0b00000000111111110000000011111111)
	val = (val>>16)&0b00000000000000001111111111111111 +
		(val & 0b00000000000000001111111111111111)
	return val
}

// NumberOfTrailingZeros returns the number of zero bits following the lowest-order ("rightmost")
//
//	one-bit in the two's complement binary representation of the specified
//	value.  Returns 64 if the specified value has no
//	one-bits in its two's complement representation, in other words if it is
//	equal to zero.
func NumberOfTrailingZeros(val int) int {
	mask32 := 0b11111111111111111111111111111111
	if val == 0 {
		return 64
	}
	n := 63
	y := val & mask32
	x := 0
	if y != 0 {
		n = n - 32
		x = y
	} else {
		x = val >> 32
	}
	y = (x << 16) & mask32
	if y != 0 {
		n = n - 16
		x = y
	}
	y = (x << 8) & mask32
	if y != 0 {
		n = n - 8
		x = y
	}
	y = (x << 4) & mask32
	if y != 0 {
		n = n - 4
		x = y
	}
	y = (x << 2) & mask32
	if y != 0 {
		n = n - 2
		x = y
	}
	return n - (((x << 1) & mask32) >> 31)
}

// HighestOneBit returns a value with at most a single one-bit, in the
//
//	position of the highest-order ("leftmost") one-bit in the specified
//	value.  Returns zero if the specified value has no
//	one-bits in its two's complement binary representation, that is, if it
//	is equal to zero.
func HighestOneBit(val int) int {
	val |= val >> 1
	val |= val >> 2
	val |= val >> 4
	val |= val >> 8
	val |= val >> 16
	val |= val >> 32
	return val - (val >> 1)
}
