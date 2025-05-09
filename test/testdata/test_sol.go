package testdata

import (
	"math/rand"
	"strconv"
	"time"
)

func GetRandomSolCuriosity() string {
	landingDate := time.Date(2012, 8, 6, 0, 0, 0, 0, time.UTC)
	now := time.Now().UTC()

	const solDuration = 88775.244 // seconds in a Martian sol

	elapsedSeconds := now.Sub(landingDate).Seconds()
	currentSol := int(elapsedSeconds / solDuration)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomSol := r.Intn(currentSol) + 1

	return strconv.Itoa(randomSol)
}
