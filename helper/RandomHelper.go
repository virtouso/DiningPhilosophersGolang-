package helper

import (
	"math/rand"
	"time"
)

func GenerateRandTime(minValue, maxValue int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(maxValue-minValue+1) + minValue)
}
