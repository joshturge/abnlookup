package abnlookup

import (
	"fmt"
	"regexp"
	"strconv"
)

// ABN and ACN weights used to calculate weighting sums
var (
	acnWeights = []int{8, 7, 6, 5, 4, 3, 2, 1}
	abnWeights = []int{10, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
)

// ValidateABN will check if an ABN is valid.
// For more information on how this works you can
// refer to: http://mathgen.ch/codes/abn.html
func ValidateABN(abn string) bool {
	// Remove all non-integer characters from the ABN
	abn = cleanNumericString(abn)

	if abn[:1] == "0" {
		return false
	}
	// Subtract 1 from the first check digit of the abn
	abnByte := []byte(abn)
	abnByte[0]--
	abn = string(abnByte)

	if abn == "" || len(abn) != 11 {
		return false
	}

	// If the weightedSum is a multiple of 89 then it's a valid ABN
	return ((calcWeightingSum("ABN", abn) % 89) == 0)
}

// ValidateACN will check if an ACN is valid.
func ValidateACN(acn string) bool {
	// Remove all non-integer characters from the ACN
	acn = cleanNumericString(acn)

	if acn == "" || len(acn) != 9 {
		return false
	}

	acnWeightingSum := calcWeightingSum("ACN", acn)

	// Convert the given check to int
	givenCheck, err := strconv.Atoi(acn[7:8])
	if err != nil {
		// Something went horribly wrong if we reach an error here
		panic(fmt.Sprintf("string must contain a non-numeric value: %s error: %s", acn[7:8], err.Error()))
	}

	// Calculate the check digit
	calcCheck := (10 - (acnWeightingSum % 10)) % 10

	return calcCheck == givenCheck

}

func cleanNumericString(str string) string {
	// Make a regexp that only allows numbers
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		// Something went horribly wrong if we reach an error here
		panic(fmt.Sprintf("couldn't compile regex: %s\n", err.Error()))
	}

	// Replace all non-integer characters with an empty string
	return reg.ReplaceAllString(str, "")
}

func calcWeightingSum(validationType string, str string) int {
	var weights []int
	switch validationType {
	case "ACN":
		weights = acnWeights
	case "ABN":
		weights = abnWeights
	}

	var weightingSum int
	var num int
	var err error
	for i := 0; i <= len(weights)-1; i++ {
		num, err = strconv.Atoi(string(str[i]))
		if err != nil {
			// Something went horribly wrong if we reach an error here
			panic(fmt.Sprintf("string must contain a non-numeric value: %s error: %s", str, err.Error()))
		}
		weightingSum += num * weights[i]
	}

	return weightingSum
}
