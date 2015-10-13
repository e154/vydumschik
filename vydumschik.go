package vydumschik

import (
	"time"
	"math/rand"
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func randomGender() string {
	gender := ""

	if i := random(0, 1); i == 1 {
		gender = "male"
	} else {
		gender = "female"
	}

	return gender
}
