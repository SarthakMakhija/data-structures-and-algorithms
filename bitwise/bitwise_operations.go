package bitwise

func IsKthBitSet(k int, number int) bool {
	numberWithKThBitSet := PowerOf2By(k - 1) //2^(k-1), power of 2 with k-1
	return number&numberWithKThBitSet != 0
}

//IsKthBitSetAlternate
//use right shift to get the bit to the first position
func IsKthBitSetAlternate(k int, number int) bool {
	return (number >> (k - 1) & 1) != 0
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

func SetBitsCountAlternateIn(number int) int {
	bitSetCount := 0
	for number > 0 {
		bitSetCount = bitSetCount + (number & 1)
		number = number >> 1
	}
	return bitSetCount
}

//SetBitsCountUsingBrianKerningam
//When you subtract 1 from a number, all the bits to the right of earliest bit set 1 are reset. Eg; 14(0001110) - 1 = 13(0001101). This algo is like unsetting the set bit
func SetBitsCountUsingBrianKerningam(number int) int {
	bitSetCount := 0
	for number > 0 {
		number = number & (number - 1)
		bitSetCount = bitSetCount + 1
	}
	return bitSetCount
}

func SetBitsCountUsingLookup(number int) int {
	bitCountLookUp := make([]int, 256)
	bitCountLookUp[0] = 0
	for index := 1; index <= 255; index++ {
		bitCountLookUp[index] = (index & 1) + bitCountLookUp[index/2]
	}

	bitSetCount := bitCountLookUp[number&0xff]
	number = number >> 8

	bitSetCount = bitSetCount + bitCountLookUp[number&0xff]
	number = number >> 8

	bitSetCount = bitSetCount + bitCountLookUp[number&0xff]
	number = number >> 8

	bitSetCount = bitSetCount + bitCountLookUp[number&0xff]
	number = number >> 8

	return bitSetCount
}

//IsPowerOf2
//Option1, Numbers which are power of 2 have only 1 bit set.
//Option2, Perform an and operation with number - 1, we should get a zero
func IsPowerOf2(number int) bool {
	if number == 0 {
		return false
	} else {
		return number&(number-1) == 0
	}
}
