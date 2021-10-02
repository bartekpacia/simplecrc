package main

func CalculateChecksum(data []byte) (crc byte) {
	crc = data[0]
	for i := 1; i < len(data)-1; i++ {
		crc ^= data[i]
	}

	return
}
