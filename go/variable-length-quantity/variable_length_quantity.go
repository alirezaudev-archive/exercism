package variablelengthquantity

import "errors"

func EncodeVarint(input []uint32) []byte {
	var result []byte

	for _, num := range input {
		encoded := encodeNumber(num)
		result = append(result, encoded...)
	}

	return result
}

func encodeNumber(num uint32) []byte {
	if num == 0 {
		return []byte{0}
	}

	var bytes []byte

	for num > 0 {
		b := byte(num & 0x7F)
		num >>= 7
		bytes = append([]byte{b}, bytes...)
	}

	for i := 0; i < len(bytes)-1; i++ {
		bytes[i] |= 0x80
	}

	return bytes
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var result []uint32
	var current uint32

	for i, b := range input {
		if current > (0xFFFFFFFF >> 7) {
			return nil, errors.New("overflow")
		}

		current = (current << 7) | uint32(b&0x7F)

		if b&0x80 == 0 {
			result = append(result, current)
			current = 0
		} else if i == len(input)-1 {
			return nil, errors.New("incomplete sequence")
		}
	}

	return result, nil
}
