package utils

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	idMutex sync.Mutex
	usedIDs = make(map[string]bool)
	idRand  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenerateUUID() string {
	return Generate3DigitID()
}

func Generate3DigitID() string {
	idMutex.Lock()
	defer idMutex.Unlock()

	// Try to find an unused ID
	for attempts := 0; attempts < 1000; attempts++ {
		// Generate random number between 100-999
		num := 100 + idRand.Intn(900)
		id := strconv.Itoa(num)

		if !usedIDs[id] {
			usedIDs[id] = true
			return id
		}
	}

	// Fallback if all IDs are used
	return strconv.Itoa(100 + idRand.Intn(900))
}
