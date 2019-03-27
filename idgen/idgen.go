package idgen

import (
	"strconv"
)

// GenerateSamID generates a SAMID for a given userID from the db
func GenerateSamID(id int) string {
	// prefix to be appended in the beginning
	prefix := "VAS19"

	zeros := ""

	switch {
	case id >= 0 && id < 10:
		zeros = "000"

	case id >= 10 && id < 100:
		zeros = "00"

	case id >= 100 && id < 1000:
		zeros = "0"

	case id >= 1000:
		zeros = ""
	}

	finalValue := zeros + strconv.Itoa(id)

	finalID := prefix + finalValue

	return finalID
}
