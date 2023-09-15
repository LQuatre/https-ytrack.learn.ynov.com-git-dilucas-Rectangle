package main

import (
	"fmt"
	"os"
	"strconv"
)

var RgbTable = map[string]string{
	"0": "\033[31m",
	"1": "\033[32m",
	"2": "\033[33m",
	"3": "\033[34m",
	"4": "\033[35m",
	"5": "\033[36m",
	"6": "\033[37m",
}

const colorNone = "\033[0m"

var Ok bool = true

func RemoveTen(count int) int {
	if count >= 10 {
		count = count - 10
		if count >= 10 {
			count = RemoveTen(count)
		}
	}
	return count
}

func main() {
	fmt.Println("Select the color :")
	for i := 0; i < 7; i++ {
		fmt.Fprintf(os.Stdout, "%s%s%s%s", RgbTable[strconv.Itoa(i)], strconv.Itoa(i), " ", colorNone)
	}

	var input string
	fmt.Scanln(&input)
	SelectedColor, err := strconv.Atoi(input)
	if err != nil {
		SelectedColor = 6
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s%s", RgbTable["0"], "Error: ", "\033[0m", "Invalid color, we will use the default color : ", strconv.Itoa(SelectedColor), "\n")
	}

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "%s%s%s%s", RgbTable["0"], "Error: ", "\033[0m", "Incomplete\n")
		return
	}

	columns, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s%s", RgbTable["0"], "Error: ", "\033[0m", "Invalid column : ", os.Args[1], "\n")
		return
	}

	raws, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s%s", RgbTable["0"], "Error: ", "\033[0m", "Invalid raws : ", os.Args[2], "\n")
		return
	}

	if raws < 0 || columns < 0 {
		fmt.Fprintf(os.Stderr, "%s%s%s%s", RgbTable["0"], "Error: ", "\033[0m", "Invalid negative argument\n")
		return
	}

	if raws == 0 || columns == 0 {
		fmt.Fprintf(os.Stderr, "%s%s%s%s", RgbTable["0"], "Error: ", "\033[0m", "not Possinble...\n")
		return
	}

	count := 0
	for i := 0; i < raws; i++ {
		for j := 0; j < columns; j++ {
			if (i == 0 && j == 0) || (i == 0 && j == columns-1) || (i == raws-1 && j == 0) || (i == raws-1 && j == columns-1) {
				fmt.Print("X")
			}
			if (i == 0 || i == raws-1) && j != 0 && j != columns-1 {
				fmt.Print("-")
			}
			if i != 0 && i != raws-1 && (j == 0 || j == columns-1) {
				printed := "0"
				if os.Args[1] != "1" {
					if Ok {
						Ok = false
						printed = strconv.Itoa(RemoveTen(count))
						count++
						if count == 10 {
							count = 0
						}
					} else {
						Ok = true
						printed = strconv.Itoa(RemoveTen(count + (raws - 3)))
					}
				} else {
					printed = strconv.Itoa(RemoveTen(count))
					count++
					if count == 10 {
						count = 0
					}
				}
				colorToPrint := RgbTable[strconv.Itoa(SelectedColor)]
				fmt.Fprintf(os.Stdout, "%s%s%s", colorToPrint, printed, colorNone)
			} else if i != 0 && i != raws-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
