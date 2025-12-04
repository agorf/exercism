package variablelengthquantity

import "errors"

func EncodeInt(num uint32) []byte {
	var bytes []byte
	for {
		b := byte(num) & 0x7F // Keep 7 right-most digits
		if len(bytes) > 0 {   // If byte is not the first
			b |= 0x80 // Set left-most bit to 1
		}
		bytes = append([]byte{b}, bytes...)
		num = num >> 7 // Discard 7 right-most digits since they have been processed
		if num == 0 {
			break
		}
	}
	return bytes
}

func EncodeVarint(input []uint32) []byte {
	var bytes []byte
	for _, num := range input {
		bytes = append(bytes, EncodeInt(num)...)
	}
	return bytes
}

func DecodeVarint(input []byte) (nums []uint32, err error) {
	var num uint32
	n := len(input)
	for i, b := range input {
		num |= (0x7F & uint32(b))
		if b&0x80 == 0 { // Last number byte if left-most bit not set
			nums = append(nums, num)
			num = 0
			continue
		}
		if i < n-1 { // Not last input byte
			num = num << 7
			continue
		}
		// Last input byte while not last number byte
		err = errors.New("incomplete sequence")
		break
	}
	return nums, err
}
