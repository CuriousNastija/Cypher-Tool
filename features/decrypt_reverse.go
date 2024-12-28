package cypher

import (
	"fmt"
	"strings"
)

func decrypt_reverse(s string) string { //trims the whitespaces of input string

s_trimmed := strings.TrimSpace(s) 

char := ""

for _, char := range s_trimmed {    		// loop that goes through input string (s_trimmed) character by character (char)
    if char >= 'A' && char <= 'Z' { 		// if capital letter A-Z
        reverse := 'Z' - (char - 'A')	 	// counts the reverse letter
        fmt.Print(string(reverse)) 			// prints reverse character (decryption)
    } else if char >= 'a' && char <= 'z' { 	// if a lowercase letter a-z
        reverse := 'z' - (char - 'a') 		// reverse character decryption
        fmt.Print(string(reverse))			// prints
    } else {
        fmt.Print(string(char))				// if not a letter, prints as is
    }

}

return char

}

func main() {
	result :=decrypt_reverse("   Svool!! Dliow?!! :)  ")
	fmt.Println(result)
}