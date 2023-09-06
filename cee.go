package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

// Crockford Base32 alphabet
const base32Alphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// Map characters to smileys
var smileyMap = map[byte]string{
	'0': "😀", '1': "😁", '2': "😂", '3': "🤣", '4': "😃",
	'5': "😄", '6': "😅", '7': "😆", '8': "😇", '9': "😈",
	'A': "😉", 'B': "😊", 'C': "😋", 'D': "😌", 'E': "😍",
	'F': "😎", 'G': "😏", 'H': "😐", 'J': "😑", 'K': "😒",
	'M': "😓", 'N': "😔", 'P': "😕", 'Q': "😖", 'R': "😗",
	'S': "😘", 'T': "😙", 'V': "😚", 'W': "😛", 'X': "😜",
	'Y': "😝", 'Z': "😞",
}

// Reverse mapping for decoding
var reverseSmileyMap = make(map[string]byte)

func init() {
	for k, v := range smileyMap {
		reverseSmileyMap[v] = k
	}
}

func encodeCrockfordBase32(data string) string {
	encodedData := ""
	smileysOnLine := 0

	for i := 0; i < len(data); {
		char := data[i]
		if strings.ContainsRune(base32Alphabet, rune(char)) {
			encodedData += smileyMap[char]
			smileysOnLine++
			if smileysOnLine == 32 {
				encodedData += "\n"
				smileysOnLine = 0
			}
			i++
		} else {
			// Handle multi-byte Unicode characters
			r, size := utf8.DecodeRuneInString(data[i:])
			if r != utf8.RuneError && size > 0 {
				encodedData += data[i : i+size]
				i += size
			} else {
				encodedData += string(char)
				i++
			}
		}
	}
	return encodedData
}

func decodeCrockfordBase32(encodedData string) string {
	decodedData := ""
	currentLine := ""

	for i := 0; i < len(encodedData); {
		char := encodedData[i]
		if strings.Contains(smileyMapString(), string(char)) {
			currentLine += string(char)
			i++
		} else {
			// Handle multi-byte Unicode characters
			r, size := utf8.DecodeRuneInString(encodedData[i:])
			if r != utf8.RuneError && size > 0 {
				currentLine += encodedData[i : i+size]
				i += size
			} else {
				currentLine += string(char)
				i++
			}
		}

		if len(currentLine) == 32 {
			decodedData += decodeSmileyLine(currentLine)
			currentLine = ""
		}
	}

	// Decode any remaining characters
	decodedData += decodeSmileyLine(currentLine)

	return decodedData
}

func decodeSmileyLine(line string) string {
	decodedLine := ""
	for _, char := range line {
		if decodedChar, found := reverseSmileyMap[string(char)]; found {
			decodedLine += string(decodedChar)
		} else {
			decodedLine += string(char)
		}
	}
	return decodedLine
}

func smileyMapString() string {
	var s strings.Builder
	for char := range smileyMap {
		s.WriteString(string(char))
	}
	return s.String()
}

func main() {
	decodeFlag := flag.Bool("d", false, "Decode using smiley encoding")
	flag.Parse()

	if *decodeFlag {
		// Decoding mode
		decodedData, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}

		decodedText := decodeCrockfordBase32(string(decodedData))
		fmt.Print(decodedText)
	} else {
		// Encoding mode
		inputData, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}

		encodedText := encodeCrockfordBase32(string(inputData))
		fmt.Print(encodedText)
	}
}

