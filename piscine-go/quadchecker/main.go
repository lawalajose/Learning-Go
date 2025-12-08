package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// readInput reads all lines from stdin and returns them as a slice
func readInput() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// getDimensions extracts width and height from the input
func getDimensions(lines []string) (int, int) {
	if len(lines) == 0 {
		return 0, 0
	}

	height := len(lines)
	width := 0

	// Use the length of the first line as width
	if height > 0 {
		width = len(lines[0])
	}

	// Validate that all lines have the same width
	for _, line := range lines {
		if len(line) != width {
			return 0, 0 // Invalid quad - inconsistent width
		}
	}

	return width, height
}

// runQuad executes a quad program with given dimensions
// Returns the output as a slice of strings, or nil if execution fails
func runQuad(quadName string, width, height int) []string {
	// Build command: ./quadX width height
	cmd := exec.Command("./"+quadName, strconv.Itoa(width), strconv.Itoa(height))

	// Capture output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Execute the command
	err := cmd.Run()
	if err != nil {
		// If the quad executable doesn't exist or fails, return nil
		return nil
	}

	// Split output into lines
	output := out.String()
	if output == "" {
		return []string{}
	}

	// Handle both Unix and Windows line endings
	output = strings.ReplaceAll(output, "\r\n", "\n")
	lines := strings.Split(output, "\n")

	// Remove trailing empty line if present (from final newline)
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

// comparePatterns checks if two patterns match exactly
func comparePatterns(pattern1, pattern2 []string) bool {
	if len(pattern1) != len(pattern2) {
		return false
	}

	for i := range pattern1 {
		if pattern1[i] != pattern2[i] {
			return false
		}
	}

	return true
}

func main() {
	// Read input from stdin
	input := readInput()

	// Handle empty input
	if len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Get dimensions from the input
	width, height := getDimensions(input)

	// Invalid dimensions
	if width == 0 || height == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Store matching quads
	var matches []string

	// Test each quad function in alphabetical order
	quadNames := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	for _, quadName := range quadNames {
		// Run the quad executable with extracted dimensions
		output := runQuad(quadName, width, height)

		// If execution succeeded and patterns match
		if output != nil && comparePatterns(input, output) {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quadName, width, height))
		}
	}

	// Output results
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}
