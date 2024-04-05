package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-WELCOME TO CIPHER METHODS*-*-*-*-*-*-*-*-*-*-*-")
	for {
		fmt.Println("Enter your option : \n1. Caesar Cipher \n2. Atbash Cipher \n3. Monoalphabetic Cipher \n4. Route Cipher \n5. Exit")
		var input int
		fmt.Scan(&input)
		switch input {
		case 1:
			cc := new(CaesarCipher)
			cc.Run()
		case 2:
			ac := new(AtbashCipher)
			ac.Run()
		case 3:
			mc := new(MonoalphabeticCipher)
			mc.Run()
		case 4:
			rc := new(RouteCipher)
			rc.Run()
		case 5:
			return
		default:
			fmt.Println("Please enter a valid option")
		}
	}
}

type CaesarCipher struct{}

func (cc *CaesarCipher) Run() {
	var text string
	var key int
	fmt.Println("1.Encrypt\n2.Decrypt\n3.Exit\nEnter your choice")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Print("Enter Text:")
		text = readInput()
		fmt.Print("Enter key:")
		fmt.Scan(&key)
		cc.CaesarEncrypt(text, key)
	case 2:
		fmt.Print("Enter Text:")
		text = readInput()
		fmt.Print("Enter key:")
		fmt.Scan(&key)
		cc.CaesarDecrypt(text, key)
	default:
		fmt.Println("Choose a valid option")
	}
}

func (cc *CaesarCipher) CaesarEncrypt(plaintext string, shift int) {
	var ciphertext strings.Builder

	for _, character := range plaintext {
		if unicode.IsUpper(character) {
			shiftedChar := rune(((int(character) - 'A' + shift) % 26) + 'A')
			ciphertext.WriteRune(shiftedChar)
		} else if unicode.IsLower(character) {
			shiftedChar := rune(((int(character) - 'a' + shift) % 26) + 'a')
			ciphertext.WriteRune(shiftedChar)
		} else {
			ciphertext.WriteRune(character)
		}
	}

	fmt.Println("The Encrypted Message is :", ciphertext.String())
}

func (cc *CaesarCipher) CaesarDecrypt(ciphertext string, shift int) {
	var plaintext strings.Builder

	for _, character := range ciphertext {
		if unicode.IsUpper(character) {
			shiftedChar := rune(((int(character) - 'A' - shift + 26) % 26) + 'A')
			plaintext.WriteRune(shiftedChar)
		} else if unicode.IsLower(character) {
			shiftedChar := rune(((int(character) - 'a' - shift + 26) % 26) + 'a')
			plaintext.WriteRune(shiftedChar)
		} else {
			plaintext.WriteRune(character)
		}
	}

	fmt.Println("The Decrypted Message is :", plaintext.String())
}

type AtbashCipher struct{}

func (ac *AtbashCipher) Run() {
	var text string
	fmt.Println("1.Encrypt\n2.Decrypt\n3.Exit\nEnter your choice")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Print("Enter Text:")
		text = readInput()
		fmt.Println("The Encrypted Message is :", ac.Atbash(text))
	case 2:
		fmt.Print("Enter Text:")
		text = readInput()
		fmt.Println("The Decrypted Message is :", ac.Atbash(text))
	default:
		fmt.Println("Choose a valid option")
	}
}

func (ac *AtbashCipher) Atbash(message string) string {
	lookupTable := map[rune]rune{
		'A': 'Z', 'B': 'Y', 'C': 'X', 'D': 'W', 'E': 'V',
		'F': 'U', 'G': 'T', 'H': 'S', 'I': 'R', 'J': 'Q',
		'K': 'P', 'L': 'O', 'M': 'N', 'N': 'M', 'O': 'L',
		'P': 'K', 'Q': 'J', 'R': 'I', 'S': 'H', 'T': 'G',
		'U': 'F', 'V': 'E', 'W': 'D', 'X': 'C', 'Y': 'B',
		'Z': 'A', '0': '9', '1': '8', '2': '7', '3': '6',
		'4': '5', '5': '4', '6': '3', '7': '2', '8': '1',
		'9': '0', 'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w',
		'e': 'v', 'f': 'u', 'g': 't', 'h': 's', 'i': 'r',
		'j': 'q', 'k': 'p', 'l': 'o', 'm': 'n', 'n': 'm',
		'o': 'l', 'p': 'k', 'q': 'j', 'r': 'i', 's': 'h',
		't': 'g', 'u': 'f', 'v': 'e', 'w': 'd', 'x': 'c',
		'y': 'b', 'z': 'a', ' ': ' ',
	}

	var cipher strings.Builder

	for _, letter := range message {
		if unicode.IsLetter(letter) || unicode.IsDigit(letter) {
			encryptedChar, found := lookupTable[unicode.ToUpper(letter)]
			if found {
				cipher.WriteRune(encryptedChar)
			}
		} else {
			cipher.WriteRune(letter)
		}
	}

	return cipher.String()
}

type MonoalphabeticCipher struct{}

func (mc *MonoalphabeticCipher) Run() {
	var text string
	fmt.Println("1.Encrypt\n2.Decrypt\n3.Exit\nEnter your choice")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Print("Enter Text:")
		text = readInput()
		mc.MonoalphabeticEncryption(text)
	case 2:
		fmt.Print("Enter Text:")
		text = readInput()
		mc.MonoalphabeticDecryption(text)
	default:
		fmt.Println("Choose a valid option")
	}
}

func (mc *MonoalphabeticCipher) MonoalphabeticEncryption(message string) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	substitution := "QWERTYUIOPASDFGHJKLZXCVBNM"
	var res strings.Builder

	for _, c := range message {
		if unicode.IsLetter(c) {
			index := strings.IndexRune(alphabet, unicode.ToUpper(c))
			if index != -1 {
				substitutedChar := rune(substitution[index])
				if unicode.IsLower(c) {
					res.WriteRune(unicode.ToLower(substitutedChar))
				} else {
					res.WriteRune(substitutedChar)
				}
			}
		} else {
			res.WriteRune(c)
		}
	}

	fmt.Println("The Encrypted Message is :", res.String())
}

func (mc *MonoalphabeticCipher) MonoalphabeticDecryption(encryptedMessage string) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	substitution := "QWERTYUIOPASDFGHJKLZXCVBNM"
	var res strings.Builder

	for _, c := range encryptedMessage {
		if unicode.IsLetter(c) {
			index := strings.IndexRune(substitution, unicode.ToUpper(c))
			if index != -1 {
				originalChar := rune(alphabet[index])
				if unicode.IsLower(c) {
					res.WriteRune(unicode.ToLower(originalChar))
				} else {
					res.WriteRune(originalChar)
				}
			}
		} else {
			res.WriteRune(c)
		}
	}

	fmt.Println("The Decrypted Message is :", res.String())
}

type RouteCipher struct{}

func (rc *RouteCipher) Run() {
	var text string
	var key1, key2 int
	fmt.Println("1.Encrypt\n2.Decrypt\n3.Exit\nEnter your choice")
	reader := bufio.NewReader(os.Stdin)
	choiceInput, _ := reader.ReadString('\n')
	choiceInput = choiceInput[:len(choiceInput)-1] // Remove newline character

	choice, err := strconv.Atoi(choiceInput)
	if err != nil || (choice != 1 && choice != 2 && choice != 3) {
		fmt.Println("Please enter a valid option")
		return
	}

	switch choice {
	case 1:
		fmt.Print("Enter Text:")
		text = readInput()
		fmt.Print("Enter key1:")
		key1 = readIntInput()
		fmt.Print("Enter key2:")
		key2 = readIntInput()
		rc.Encrypt(text, key1, key2)
	case 2:
		fmt.Print("Enter Text:")
		text = readInput()
		fmt.Print("Enter Key1:")
		key1 = readIntInput()
		fmt.Print("Enter Key2:")
		key2 = readIntInput()
		rc.Decrypt(text, key1, key2)
	case 3:
		return
	default:
		fmt.Println("Choose a valid option")
	}
}

var grid [][]rune

func (rc *RouteCipher) fillGrid(message string, key1, key2 int) {
	grid = make([][]rune, key1)
	messageIndex := 0
	for col := 0; col < key2; col++ {
		for row := 0; row < key1; row++ {
			if messageIndex < len(message) {
				grid[row] = append(grid[row], rune(message[messageIndex]))
				messageIndex++
			} else {
				grid[row] = append(grid[row], ' ')
			}
		}
	}
}

func (rc *RouteCipher) Encrypt(text string, key1, key2 int) {
	rc.fillGrid(text, key1, key2)
	var cipherText strings.Builder
	for row := 0; row < key1; row++ {
		for col := 0; col < key2; col++ {
			cipherText.WriteRune(grid[row][col])
		}
	}
	fmt.Println("Encrypted Message:", cipherText.String())
}

func (rc *RouteCipher) Decrypt(text string, key1, key2 int) {
	fillGridDecrypt(text, key1, key2)
	var plainText strings.Builder
	for col := 0; col < key2; col++ {
		for row := 0; row < key1; row++ {
			plainText.WriteRune(grid[row][col])
		}
	}
	fmt.Println("Decrypted Message:", plainText.String())
}

func fillGridDecrypt(message string, key1, key2 int) {
	grid = make([][]rune, key1)
	messageIndex := 0
	for row := 0; row < key1; row++ {
		grid[row] = make([]rune, key2)
		for col := 0; col < key2; col++ {
			if messageIndex < len(message) {
				grid[row][col] = rune(message[messageIndex])
				messageIndex++
			}
		}
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readIntInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	return num
}
