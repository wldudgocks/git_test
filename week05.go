package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	input := reader.ReadString('\n')
	fmt.Println(input)
}
