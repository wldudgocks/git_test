package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n') // option 2
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = string.TrimSpace(inputScoreString) //remove space
	inputScore, err := strconv.ParseFloat(inputScoreString, 32)
	fmt.Println(inputScore)
	if inputScore >= 90 {
		grade := "A grade!"

	}
	else {
		grade := "under A grade..."
	}
}
