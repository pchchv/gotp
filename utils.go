package otp

import "time"

// integer to byte array
func Itob(integer int64) []byte {
	byteArr := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		byteArr[i] = byte(integer & 0xff)
		integer = integer >> 8
	}
	return byteArr
}

// get current timestamp
func currentTimestamp() int64 {
	return time.Now().Unix()
}
