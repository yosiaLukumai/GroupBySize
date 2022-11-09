package helpers

import (
	"math"
)

// Func return the converted size of bytes in mbs or
func SizeConverter(size int64, output, places int) float64 {
	// by default it retures the mbs
	// 1 ==> KB's 2 ==> MB's 3 ==> GB's 
	var outputSize float64
	var specificNumber float64 = 1024
	if output == 1 {
		outputSize = (float64(size) / specificNumber)
		return RoundUp(outputSize, places)
	}else if output == 2 {
		outputSize = (float64(size) / (specificNumber * specificNumber))
		return RoundUp(outputSize, places)
	}else if output == 3 {
		outputSize = float64(float64(size) / (specificNumber * specificNumber * specificNumber))
		return RoundUp(outputSize, places)
	} else {
		return RoundUp(outputSize, places)
	}
	
}

func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return newVal
}