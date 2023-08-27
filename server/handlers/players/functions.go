package players_handler

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func GeneratePassword() string {
	var charTypes = [3]string{"number", "symbol", "letter"}
	passwordLength := 15
	password := ""
	for i := 0; i <= passwordLength; i++ {
		charType := charTypes[rand.Intn(len(charTypes))]
		switch charType {
		case "number":
			password += randomNumber()
		case "symbol":
			password += randomSymbol()
		case "letter":
			password += randomLetter()
		}
	}
	fmt.Println(password)
	return password
}

func randomNumber() string {
	return strconv.Itoa(rand.Intn(10))
}

func randomLetter() string {
	alphabet := make([]string, 0)
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	// rang 0 - 26
	letter := alphabet[rand.Intn(len(alphabet) - 1)]
	// rang 0 - 1
	n := rand.Intn(2)
	var shouldBeUpperCased bool
	if n == 1 {
		shouldBeUpperCased = true
	} else {
		shouldBeUpperCased = false
	}
	if shouldBeUpperCased {
		return strings.ToUpper(letter)
	}
	return letter
}

func randomSymbol() string {
	var symbols = [5]string{"-", ".", "@", "&", "%"}
	return symbols[rand.Intn(len(symbols) - 1)]
}
