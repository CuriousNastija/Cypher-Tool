package cypher

import (
	"fmt"
	"strings"
)

func encrypt_reverse(s string) string {

s_trimmed := strings.TrimSpace(s) //trims the whitespaces of input string

char := ""

for _, char := range s_trimmed {    		// loop that goes through input string (s_trimmed) character by character (char)
    if char >= 'A' && char <= 'Z' { 		// if capital letter A-Z
        reverse := 'Z' - (char - 'A')	 	// counts the reverse letter
        fmt.Print(string(reverse)) 			// prints reverse character (encryption)
    } else if char >= 'a' && char <= 'z' { 	// if a lowercase letter a-z
        reverse := 'z' - (char - 'a') 		// reverse character encryption
        fmt.Print(string(reverse))			// prints
    } else {
        fmt.Print(string(char))				// if not a letter, prints as is
    }

}

return char

}

func main() {
	result :=encrypt_reverse("   Hello!! World?!! :)  ")
	fmt.Println(result)
}


/*
rakenna ohjelma joka ottaa vastaan merkkijonon (string), 
muuttaa jokaisen kirjaimen erikseen aakkosten vastakkaiseksi.

1 trimmaus
Before validating the input, it has to be trimmed (remove whitespaces 
from the beginning and the end of the input).

2 if lause jossa kysytään onko ulkona kirjainten unicode alueelta
When encrypting or decrypting, ensure that any non-alphabet characters 
in the message are left unchanged.

Enter the message: //input
zb

Decrypted message using reverse: //output
ay

*/