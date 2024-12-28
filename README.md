# Cypher

## Description

The purpose of this tool is to provide user a simple command-line based encryption/decryption tool. 
With the Cypher tool user can choose from three different encryption methods and either encrypt or decrypt
their message. 

---

# How to use Cypher

The program can be executed from the command line typing 'go run cypher.go'.

When the program is executed, the user is greeted with “Welcome to the Cypher Tool!” message and presented
with two options: 1) Encrypt, and 2) Decrypt. The selection is then made with pressing the corresponding
number ('1' or '2').

After making the selection, the user can then continue to select the preferred cypher method. The encyption
options are: 1) Rot13, 2) Reverse and 3) XOR. Selection is made using numbers '1', '2' or '3'.

## Using Rot13 and Reverse

After selecting the encryption method, the user can then enter the message they want to proceed with. After
typing in the message and hitting the 'Enter' key, the encypted/decrypted message is then printed on screen.

## Using XOR

After selecting the XOR encryption method, user then has to enter the preferred 'key' used to encrypt the
message. The key is applied hitting 'Enter'. Encrypted/decrypted message is then printed on the screen. The
entered 'key' is needed in order to decrypt the input message correctly afterwards.

---

# The Encryption Methods

The cypher tool provides three different encryption methods: Rot13, Reverse and XOR. 

## ROT13

In Rot13 encryption the letters A-Z  are substituted with the 13th letter after it in the Latin alphabet.
To undo the encryption the same algorithm is being applied.

## Reverse

Reverse encyption method takes each letter of the input message and returns its reverse letter in the Latin
alphabet. For example 'a' becomes 'z', 'b' becomes 'y', and so on. The encryption can be undone by using
the same algorithm.

## XOR

XOR encryption is an encryption method that combines each bit of plaintext with a corresponding bit from a
'key' selected by the user. For example, if both bits are the same (0,0 or 1,1), the result is 0; if the
bits are different (0,1 or 1,0), the result is then 1. The encryption can then be undone applying the same
'key' again. 

---

# Example usage

go run main.go

## Example output

Encrypted (hex): 1e1c001604341d111f150f371d04250c26001002
Decrypted: Hello World!


