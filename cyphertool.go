package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to the Cypher Tool!")

	for {

	toEncrypt, encoding, message := getInput() //The variables toEncrypt, encoding, and message 
	//are being assigned the respective values returned by the getInput() function.

	//loop1
	if toEncrypt { 
		switch encoding {
		case "ROT13":
			fmt.Println(encrypt_rot13(message))
		case "Reverse":
			fmt.Println(encrypt_reverse(message))
		case "XOR":
			fmt.Println(encrypt_xor(message))
		default:
			fmt.Println("Invalid encoding method.")
			return
		}
	} else {
		switch encoding {
		case "ROT13":
			fmt.Println(decrypt_rot13(message))
		case "Reverse":
			fmt.Println(decrypt_reverse(message))
		case "XOR":
			fmt.Println(decrypt_xor(message))
		default:
			fmt.Println("Invalid encoding method.")
			return
		}
	}
}
}

func getInput() (toEncrypt bool, encoding string, message string) {

	
	var operation string
//loop2
	for { // keeps running until user provides a valid input
	fmt.Println("Select operation (1/2):")
	fmt.Println("1. Encrypt.")
	fmt.Println("2. Decrypt.")
	fmt.Scan(&operation)

	if operation != "1" && operation != "2" {
		fmt.Println("Invalid operation selected. Try again.")
		continue // restart the loop
	}
	break // if input valid exit the loop
}

	toEncrypt = (operation == "1") // assignment with comparison, if = 1, then true

	//loop3
	for {
	fmt.Println("Select cypher (1/2/3):")
	fmt.Println("1. ROT13.")
	fmt.Println("2. Reverse.")
	fmt.Println("3. XOR.")
	fmt.Scan(&encoding)

	switch encoding {
	case "1":
		encoding = "ROT13"
		break
	case "2":
		encoding = "Reverse"
		break
	case "3":
		encoding = "XOR"
		break
	default:
		fmt.Println("Invalid cypher selected. Try again.")
		continue // restarts the lopp until valid answer
	}
	break
	}


	reader := bufio.NewReader(os.Stdin) // more flexible reader that can accept spaces, buffered reader, which reads from standard input = console
	fmt.Printf("Enter the message: ")
	message, _ = reader.ReadString('\n') // reads until Enter
	message = strings.TrimSpace(message) // removes spaces

	return toEncrypt, encoding, message
}

func encrypt_rot13(s string) string {

	result := ""

	for _, char := range s { // loop through each character
		if char >= 'A' && char <= 'Z' { // if capital letter
			encoded := 'A' + (char-'A'+13)%26 // Apply ROT13 for capital letters 'A' + convert shifted position back to character
			result += string(encoded)         // append the encoded character to result
		} else if char >= 'a' && char <= 'z' { // if lowercase letter
			encoded := 'a' + (char-'a'+13)%26 // Apply ROT13 for lowercase letters
			result += string(encoded)         // append the encoded character to result
		} else {
			result += string(char) // if not a letter, append as is
		}
	}

	return result
}

func encrypt_reverse(s string) string {
	result := ""

	for _, char := range s { // loop through each character
		if char >= 'A' && char <= 'Z' { // if capital letter
			reverse := 'Z' - (char - 'A') // reverse the alphabet
			result += string(reverse)     // append the reverse character to result
		} else if char >= 'a' && char <= 'z' { // if lowercase letter
			reverse := 'z' - (char - 'a') // reverse the alphabet
			result += string(reverse)     // append the reverse character to result
		} else {
			result += string(char) // if not a letter, append as is
		}
	}

	return result
}

func decrypt_rot13(s string) string {

	result := ""

	for _, char := range s { // loop through each character
		if char >= 'A' && char <= 'Z' { // if capital letter
			decoded := 'A' + (char-'A'+13)%26 // Apply ROT13 for capital letters
			result += string(decoded)         // append the decoded character to result
		} else if char >= 'a' && char <= 'z' { // if lowercase letter
			decoded := 'a' + (char-'a'+13)%26 // Apply ROT13 for lowercase letters
			result += string(decoded)         // append the decoded character to result
		} else {
			result += string(char) // if not a letter, append as is
		}
	}

	return result
}

func decrypt_reverse(s string) string {
	result := ""

	for _, char := range s { // loop through each character
		if char >= 'A' && char <= 'Z' { // if capital letter
			reverse := 'Z' - (char - 'A') // reverse the alphabet
			result += string(reverse)     // append the reverse character to result
		} else if char >= 'a' && char <= 'z' { // if lowercase letter
			reverse := 'z' - (char - 'a') // reverse the alphabet
			result += string(reverse)     // append the reverse character to result
		} else {
			result += string(char) // if not a letter, append as is
		}
	}

	return result
}

func encrypt_xor(s string) string {

	reader := bufio.NewReader(os.Stdin)

	var key string
	for { 											//loop 4
	fmt.Printf("Please, enter a key to encrypt: ")
	keyInput, _ := reader.ReadString('\n') // read input until new line, ignore error
	keyInput = strings.TrimSpace(keyInput)      // trim leading or trailing new lines or spaces

	//check key length
	if len(keyInput) > 0 {
		key = keyInput
		break //exit loop if key is not empty
	}
	fmt.Println("The key field is empty. Please, try again.")
}

	// convert s into a slice of bytes
	// each element stores decimal ASCII value of a character in s
	byteArrayS := []byte(s)

	// if key length is smaller than s length, then key is repeated until it matches the s length
	if len(key) < len(s) {
		for i := len(key); i < len(s); i++ {
			key += string(key[i%len(key)]) // modulo operation to wrap around key and convert to string
		}
	}

	// if key is longer then truncate the key
	if len(key) > len(s) {
		key = key[:len(s)]
	}

	// convert key into a slice of bytes
	byteArrayKey := []byte(key)
	// create a byte slice to store the result of XOR operation
	encryptedBytes := make([]byte, len(byteArrayS))

	// XOR operation bit by bit converted to decimal values of bytes stored in the slice encryptedBytes
	for i := 0; i < len(byteArrayS); i++ {
		encryptedBytes[i] = byteArrayS[i] ^ byteArrayKey[i]
	}

	// convert encrypted bytes to hex using func bytesToHex
	return bytesToHex(encryptedBytes)

}

// decryption function
//loop 5
func decrypt_xor(encryptedS string) string {

	reader := bufio.NewReader(os.Stdin)

	var key string
	for {
	fmt.Printf("Please, enter the key to decrypt: ")
	keyInput, _ := reader.ReadString('\n') // read input until new line, ignore error
	keyInput = strings.TrimSpace(keyInput)      // trim leading or trailing new lines or spaces

	// check key length
	if len(keyInput) > 0 {
		key = keyInput
		break
	}
		fmt.Println("The key field is empty. Please, try again.") 
	}
	

	// decode hex string back to byte array using func hexToBytes
	encryptedBytes := hexToBytes(encryptedS)

	if len(encryptedBytes) == 0 {
		// Invalid input, restart program
		fmt.Println("Invalid encrypted message. Restarting input process.")
		return "" // Return empty to signal failure, allowing restart in main loop
	}

	if len(key) < len(encryptedBytes) {
		for i := len(key); i < len(encryptedBytes); i++ {
			key += string(key[i%len(key)])
		}

	}

	if len(key) > len(encryptedBytes) {
		key = key[:len(encryptedBytes)]
	}

	// convert key into a slice of bytes
	byteArrayKey := []byte(key)
	// create a byte slice to store decrypted result
	decryptedBytes := make([]byte, len(encryptedBytes))

	// XOR operation for decription
	for i := 0; i < len(encryptedBytes); i++ {
		decryptedBytes[i] = encryptedBytes[i] ^ byteArrayKey[i]
	}

	return string(decryptedBytes)

}

// function to convert byte array into hexadecimal string
func bytesToHex(byteArray []byte) string {

	hexString := "" // initialize empty string to store result

	// loop iterates through each element in byteArray slice
	// _ ignores the index of the element in array
	// b variable representing value of each byte in array
	for _, b := range byteArray {
		// formatted string conversion converts each byte into 2- character (02) hexadecimal (x) string
		hexString += fmt.Sprintf("%02x", b)
	}

	return hexString
}

// function to convert hexadecimal string into array of bytes
func hexToBytes(hexString string) []byte {
	
	if len(hexString) < 2 {
		// If hex string is not of even length, it may be corrupted or incomplete
		fmt.Println("Your message is too short.")
		//getInput()
		return []byte{}
	}
	
	// hex uses 2 characters to represent each byte, so the length of array should be 2-times smaller
	byteArray := make([]byte, len(hexString)/2)

	// loops through string with step 2, because each byte is represented by 2 hex characters
	for i := 0; i < len(hexString); i += 2 {
		// converts pair of hex characters into 1 byte
		// hexString[i:i+2] takes substring of hexString starting at i and takes the next to characters
		// strconv.ParseUnit converts string representing a number into unsigned integer (unit64), number in 16 - base, result will be in 8-bit
		// _ ignores second return value - error
		byteValue, _ := strconv.ParseUint(hexString[i:i+2], 16, 8)
		// stores resulting byteValue in the appropriate place in array
		// since we incremented by step 2, the index will be half of i
		byteArray[i/2] = byte(byteValue)
	}

	return byteArray
}
