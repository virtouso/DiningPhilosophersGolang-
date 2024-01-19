package main

import (
	"fmt"
	"github.com/virtouso/DiningPhilosophersGolang-/models"
	"strconv"
)

var userInput string
var philosophersCount int
var inputError error

var philosophers []models.Philosopher
var chopsticks []models.Chopstick

func main() {

	for philosophersCount <= 2 || philosophersCount >= 10 {
		fmt.Println("write number of dining philosophers (any integer value 3-9. small number for testability. big numbers also tested):")
		fmt.Scanln(&userInput)
		philosophersCount, inputError = strconv.Atoi(userInput)
		if inputError != nil {
			fmt.Println("invalid input. try again.")
		}
	}

	println("starting the infinite eat process")

	philosophers = make([]models.Philosopher, philosophersCount)
	chopsticks = make([]models.Chopstick, philosophersCount)

}
