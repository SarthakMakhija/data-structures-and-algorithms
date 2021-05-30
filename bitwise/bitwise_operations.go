package bitwise

func IsKthBitSet(k int, number int) bool {
	numberWithKThBitSet := PowerOf2By(k - 1) //2^(k-1), power of 2 with k-1
	return number&numberWithKThBitSet == numberWithKThBitSet
}

func PowerOf2By(k int) int {
	base := 1
	return base << k
}

func SetBitsCountIn(number int) int {
	const intSizeInBits = 32
	base := 1
	bitSetCount := 0

	for count := 1; count <= intSizeInBits; count++ {
		if number&base >= 1 {
			bitSetCount = bitSetCount + 1
		}
		base = base << 1
	}
	return bitSetCount
}
