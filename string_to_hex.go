package ibus

func stringToBytes(str string) []byte {
	bytes := make([]byte, 0)
	chars := []rune(str)
	for _, el := range chars {
		bytes = append(bytes, byte(int(el)))
	}
	return bytes
}
