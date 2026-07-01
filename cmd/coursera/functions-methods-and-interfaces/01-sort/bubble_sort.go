package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}

func BubbleSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}

	// Standard bubble sort: after each pass, the largest element "bubbles" to the end.
	for end := n - 1; end > 0; end-- {
		swapped := false
		for i := 0; i < end; i++ {
			if s[i] > s[i+1] {
				Swap(s, i)
				swapped = true
			}
		}
		if !swapped {
			return // already sorted
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter up to 10 integers (space-separated): ")

	line, err := reader.ReadString('\n')
	if err != nil {
		// still try to parse whatever we got (e.g., no trailing newline)
	}
	line = strings.TrimSpace(line)
	if line == "" {
		fmt.Println("No input provided.")
		return
	}

	parts := strings.Fields(line)
	if len(parts) > 10 {
		fmt.Println("Please enter no more than 10 integers.")
		return
	}

	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		v, err := strconv.Atoi(p)
		if err != nil {
			fmt.Printf("Invalid integer: %q\n", p)
			return
		}
		nums = append(nums, v)
	}

	BubbleSort(nums)

	// Print sorted list on one line
	for i, v := range nums {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
