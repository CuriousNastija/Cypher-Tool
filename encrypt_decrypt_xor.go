package cypher

/* Encryption using XOR (exclusive OR). This is a symmetric encryption method with a key.
Each character of the message (represented as its ASCII value) undergoes XOR operation with a corresponding character of the key.
As the result we have a byte array, where each element represents XOR-ed value:
 - if bits are different, XOR returns 1
 - if bits are the same, XOR returns 0
Not all ASCII characters are printable, so converted to string the encrypted message might look empty or weird.
Therefore, for readability the byte array is converted into a hexadecimal string.
During decryption procedure this hex string is converted back into original byte array.
The same key is used to perform XOR operation again, decrypting it into original message. */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// encryption function
func encrypt_xor(s string) string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please, enter a key to encrypt: ")
	key, _ := reader.ReadString('\n') // read input until new line, ignore error
	key = strings.TrimSpace(key)  // trim leading or trailing new lines or spaces

	
	// convert s into a slice of bytes
	// each element stores byte value of UTF-8 encoded value of a character in s
	byteArrayS := []byte(s) 
	
	// check key length
	// if key is empty > return
	if len(key) == 0 {
		fmt.Println("The key field is empty. Please, try again.")
		return ""
	}
	
	// if key is longer then truncate the key
	if len(key) > len(s) {
		key = key[:len(s)]
	}

	// if key length is smaller than s length, then key is repeated until it matches the s length
	if len(key) < len(s) {
		for i := len(key); i < len(s); i++ {
			key += string(key[i % len(key)]) // modulo operation to wrap around key
		}
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
func decrypt_xor(encryptedS string) string {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please, enter the key to decrypt: ")
	key, _ := reader.ReadString('\n') // read input until new line, ignore error
	key = strings.TrimSpace(key)  // trim leading or trailing new lines or spaces
	
	// decode hex string back to byte array using func hexToBytes
	encryptedBytes := hexToBytes(encryptedS)

	// check key length
	if len(key) == 0 {
		fmt.Println("The key field is empty. Please, try again.")
		return ""
	}

	// if key is longer then truncate the key
	if len(key) > len(s) {
		key = key[:len(s)]
	}

	if len(key) < len(encryptedBytes) {
		for i := len(key); i < len(encryptedBytes); i++ {
			key += string(key[i % len(key)])
		}

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
func bytesToHex(byteArray[] byte) string {

	hexString := "" // initialize empty string to store result

	// loop iterates through each element in byteArray slice
	// _ ignores the index of the element in array
	// b variable representing value of each bite in array
	for _, b := range byteArray {
		// formatted string conversion converts each byte into 2- character (02) hexadecimal (x) string
		hexString += fmt.Sprintf("%02x", b) 
	}

	return hexString
}

// function to convert hexadecimal string into array of bytes
func hexToBytes(hexString string) []byte {
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


/* func main() {

	message := "Good5!nklskj****"

	encrypted := encrypt_xor(message)
	fmt.Println("Encrypted (hex):", encrypted)

	decrypted := decrypt_xor(encrypted)
	fmt.Println("Decrypted:", decrypted)

	
} */