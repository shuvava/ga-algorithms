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
