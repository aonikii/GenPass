package password

import (
	"crypto/rand"
	"log"
	"math/big"
	"strconv"
	"strings"
)

const (
	letters                                  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits                                   = "0123456789"
	specialSymbols                           = "!$^&*()-=+[]{}"
	minDigitsAmountForLengthLessThanEleven   = 2
	minDigitsAmountForLengthMoreThanTen      = 3
	minSpecSymbAmountForLengthLessThanEleven = 1
	minSpecSymbAmountForLengthMoreThanTen    = 2
)

func GeneratePass(length, isSpecSym string) string {
	var password strings.Builder
	symbPositions := positions(length, isSpecSym)
	digitsSlice := symbPositions["d"]
	specSymbSlice := symbPositions["s"]
	lengthInInt, err := strconv.Atoi(length)
	if err != nil {
		log.Panic(err)
	}

	lettersAmount := lengthInInt - len(digitsSlice) - len(specSymbSlice)
	setOfLetters, ok := chooseLetters(lettersAmount)
	counterLetters := 0
	for !ok {
		setOfLetters, ok = chooseLetters(lettersAmount)
	}
	for i := range lengthInInt {
		if contains(digitsSlice, i) {
			add, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
			if err != nil {
				log.Panic(err)
			}
			password.WriteString(string(digits[int(add.Int64())]))
		} else if len(specSymbSlice) != 0 && contains(specSymbSlice, i) {
			add, err := rand.Int(rand.Reader, big.NewInt(int64(len(specialSymbols))))
			if err != nil {
				log.Panic(err)
			}
			password.WriteString(string(specialSymbols[int(add.Int64())]))
		} else {
			password.WriteString(string(setOfLetters[counterLetters]))
			counterLetters++
		}
	}
	return password.String()
}

func numAmount(l string) int {
	length, err := strconv.Atoi(l)
	if err != nil {
		log.Panic(err)
	}
	digitsAmount, err := rand.Int(rand.Reader, big.NewInt(3)) //[0,1,2]
	if err != nil {
		log.Panic(err)
	}
	if length <= 10 {
		return int(digitsAmount.Int64()) + minDigitsAmountForLengthLessThanEleven
	} else {
		return int(digitsAmount.Int64()) + minDigitsAmountForLengthMoreThanTen
	}
}

func specSymbAmount(l string) int {
	length, err := strconv.Atoi(l)
	if err != nil {
		log.Panic(err)
	}
	symbAmount, err := rand.Int(rand.Reader, big.NewInt(2)) //[0,1]
	if err != nil {
		log.Panic(err)
	}
	if length <= 10 {
		return int(symbAmount.Int64()) + minSpecSymbAmountForLengthLessThanEleven
	} else {
		return int(symbAmount.Int64()) + minSpecSymbAmountForLengthMoreThanTen
	}
}

func positions(l, isSpecSym string) map[string][]int {
	length, err := strconv.Atoi(l)
	if err != nil {
		log.Panic(err)
	}
	pos := make(map[string][]int, 2)
	counter := 0
	numAmo := numAmount(l)
	for counter != numAmo {
		ranPos, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			log.Panic(err)
		}
		resCh := int(ranPos.Int64())
		if contains(pos["d"], resCh) {
			continue
		}
		pos["d"] = append(pos["d"], resCh)
		counter++
	}
	if isSpecSym == "on" {
		specSymbAmo := specSymbAmount(l)
		counter = 0
		for counter != specSymbAmo {
			ranPos, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
			if err != nil {
				log.Panic(err)
			}
			resCh := int(ranPos.Int64())
			if contains(pos["s"], resCh) || contains(pos["d"], resCh) {
				continue
			}
			pos["s"] = append(pos["s"], resCh)
			counter++
		}
	}
	return pos
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func chooseLetters(letterAmount int) (string, bool) {
	var res strings.Builder
	var ok bool
	for range letterAmount {
		ranLet, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			log.Panic(err)
		}
		resLetter := letters[int(ranLet.Int64())]
		if resLetter >= 65 && resLetter <= 90 {
			ok = true
		}
		res.WriteString(string(letters[int(ranLet.Int64())]))
	}
	return res.String(), ok
}
