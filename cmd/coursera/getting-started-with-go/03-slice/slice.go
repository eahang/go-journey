package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Program: Sorted Integer Slice Collector
//
// Prompts the user to enter integers and stores them in a sorted slice.
// The slice starts with length 0 and capacity 3, and grows as needed.
// After each entry, the slice is sorted and printed.
// Enter 'X' or 'x' to quit the program.

func main() {
	nums := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter an integer or X to quit: ")
		if !scanner.Scan() {
			fmt.Println("Error reading input. Exiting.")
			break
		}
		text := scanner.Text()
		if text == "X" || text == "x" {
			break
		}
		var value int
		_, err := fmt.Sscanf(text, "%d", &value)
		if err != nil {
			fmt.Println("Invalid input, please enter an integer.")
			continue
		}
		nums = append(nums, value)
		sort.Ints(nums)
		fmt.Println(nums)
	}
	fmt.Println("Exiting program.")
}
