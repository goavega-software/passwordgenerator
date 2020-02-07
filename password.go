package passwordgen

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

var consonants = [21]string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
var vowels = [5]string{"a", "e", "i", "o", "u"}
var leetMap = map[string]int{
	"a": 4,
	"e": 3,
	"i": 1,
	"o": 0,
}
var symbols = [6]string{"!", "#", "$", "@", "%", "^"}
var hasLeetAlready = false

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
Generate generates password
*/
func Generate(request GenerateRequest) string {
	length := 8.0
	password := ""
	numDigits := 2
	alphaNumChars := int(length) - numDigits
	pairs := int(math.Floor(float64(alphaNumChars / 2)))
	for pairs > 0 {
		password += getPair(request.LeetMode, request.RandomUpperCase)
		pairs--
	}
	password += getDigits(numDigits)
	if request.AddSymbol {
		password += symbols[getRand(len(symbols)-1)]
	}
	hasLeetAlready = false
	return password
}

func getPair(leetMode bool, randomUpperCase bool) string {
	return getConsonant() + getVowel(leetMode)
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

func getVowel(leetMode bool) string {
	var vowel = vowels[getRand(len(vowels)-1)]
	if val, ok := leetMap[vowel]; ok && leetMode && !hasLeetAlready {
		vowel = strconv.Itoa(val)
		hasLeetAlready = true
	}
	return vowel
}

func getRand(max int) int {
	return rand.Intn(max)
}
