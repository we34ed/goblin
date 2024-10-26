// Программа показывает, сдал ли пользователь экзамен

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your weight: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	bull := strings.TrimSuffix(input, "\n")
	if err != nil {
		log.Fatal(err)
	}
	weight, err := strconv.Atoi(bull)
	if err != nil {
		log.Fatal(err)
	}
	if weight >= 100 {
		fmt.Println("You are a lightweight ")
	} else if weight >= 75 {
		fmt.Println("You are sreniy ves!")
	} else {
		fmt.Println("You are a heavyweight!!! AXAXAXA!")
	}
}

//todo:конец
