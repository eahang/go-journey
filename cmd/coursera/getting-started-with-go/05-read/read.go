package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Objective:
// Write a program which reads information from a file and represents it in a slice of structs.
// Assume that there is a text file which contains a series of names.
// Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
// Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
// Each struct created will be added to a slice, and after all lines have been read from the file,
// your program will have a slice containing one struct for each line in the file. After reading all lines from the file,
// your program should iterate through your slice of structs and print the first and last names found in each struct.

type name struct {
	fname [20]byte
	lname [20]byte
}

func put20(s string) [20]byte {
	var a [20]byte
	s = strings.TrimSpace(s)
	if len(s) > 20 {
		s = s[:20]
	}
	copy(a[:], s)
	return a
}

func getString20(a [20]byte) string {
	// Convert [20]byte to string, removing trailing zero bytes
	return strings.TrimRight(string(a[:]), "\x00")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter file name: ")
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading file name:", err)
		return
	}
	filename = strings.TrimSpace(filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	var names []name

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			// skip bad lines (or you could print a warning)
			continue
		}

		n := name{
			fname: put20(parts[0]),
			lname: put20(parts[1]),
		}
		names = append(names, n)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, n := range names {
		fmt.Printf("%s %s\n", getString20(n.fname), getString20(n.lname))
	}
}
