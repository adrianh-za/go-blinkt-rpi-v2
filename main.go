package main

import (
	"bufio"
	"fmt"
	"go-blinkt-rpi-v2/examples"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose an example to run (1-5):")
		fmt.Println("1: Kitt")
		fmt.Println("2: Kitt Colours")
		fmt.Println("3: Colours")
		fmt.Println("4: Rainbow")
		fmt.Println("5: Rainbow2")
		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(input[:len(input)-1])
		if err != nil || choice < 1 || choice > 5 {
			fmt.Println("Invalid choice, please choose again.")
			continue
		}

		switch choice {
		case 1:
			examples.Kitt()
		case 2:
			examples.KittColours()
		case 3:
			examples.Colours()
		case 4:
			examples.Rainbow()
		case 5:
			examples.Rainbow2()
		}
		break
	}
}