//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumNum(bytes []byte) int {
	var num string
	var storeNum []string
	var sum int

	for _, r := range bytes {
		if r != '\n' {
			num += string(r)
		} else {
			storeNum = append(storeNum, num)
			num = ""
		}
	}
	for i := 0; i < len(storeNum); i++ {
		digit, _ := strconv.Atoi(storeNum[i])
		sum += digit
	}
	return sum
}
func sumNum1(fold string) int {
	file, err := os.Open(fold)
	if err != nil {
		fmt.Printf("Error :%v ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scan error : %v", err)
	}
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	soln := sumNum1(files[3])
	fmt.Println(soln)

}
