package passwordgen

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

var consonants = [21]string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
var vowels = [5]string{"a", "e", "i", "o", "u"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Generate() string {
	length := 8.0
	password := ""
	numDigits := int(math.Floor(length * 0.5))
	alphaNumChars := int(length) - numDigits
	pairs := int(math.Floor(float64(alphaNumChars / 2)))
	for pairs > 0 {
		password += getPair()
		pairs--
	}
	password += getDigits(numDigits)
	return password
}

func getPair() string {
	return getConsonant() + getVowel()
}

func getDigits(length int) string {
	var digits = ""
	for i := 0; i < length; i++ {
		digits += getDigit()
	}
	return digits
}

func getDigit() string {
	return strconv.Itoa(getRand(9))
}

func getConsonant() string {
	return consonants[getRand(len(consonants)-1)]
}

func getVowel() string {
	return vowels[getRand(len(vowels)-1)]
}

func getRand(max int) int {
	return rand.Intn(max)
}
