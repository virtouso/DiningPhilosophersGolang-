package main

import (
	"fmt"
	"github.com/virtouso/DiningPhilosophersGolang-/helper"
	"github.com/virtouso/DiningPhilosophersGolang-/models"
	"strconv"
	"time"
)

var userInput string
var philosophersCount int
var inputError error

var philosophers []*models.Philosopher
var chopsticks []*models.Chopstick

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

	philosophers = make([]*models.Philosopher, philosophersCount)
	chopsticks = make([]*models.Chopstick, philosophersCount)

	for i := 0; i < philosophersCount; i++ {
		chopsticks[i] = &models.Chopstick{
			BeingUsed: false,
			Id:        i,
		}
	}

	for i := 0; i < philosophersCount; i++ {
		philosophers[i] = &models.Philosopher{
			Id:       i,
			EatCount: 0,
		}
	}

	println("starting the infinite eat process")

	ch1 := make(chan int, philosophersCount)
	ch2 := make(chan int, philosophersCount)

	go eat(ch1, ch2)
	go think(ch2, ch1)

	for index, _ := range philosophers {
		openChops := chopsticks[helper.MakeIndex(index, philosophersCount)].BeingUsed || chopsticks[helper.MakeIndex(index+1, philosophersCount)].BeingUsed
		if !openChops {

			ch1 <- index
		}
	}

	time.Sleep(time.Second * 20) // to make sure the simulation end after certain amount of time.
	close(ch1)
	close(ch2)

	// logging to show result of the simulation
	for _, val := range philosophers {
		println(helper.JoinStrings("philosopher:", strconv.Itoa(val.Id), "->Eat Count:", strconv.Itoa(val.EatCount)))
	}
	println("as you can see no one has starved!...")

}

func eat(ch1 <-chan int, ch2 chan<- int) {

	for {
		val, ok := <-ch1
		if !ok {
			fmt.Println("Message delivery Bug. Exiting.")
			return
		}
		//	philosophers[val].Eating = true
		philosophers[val].EatCount += 1

		canEat := !chopsticks[helper.MakeIndex(val, philosophersCount)].BeingUsed && !chopsticks[helper.MakeIndex(val+1, philosophersCount)].BeingUsed
		if canEat {

			chopsticks[helper.MakeIndex(val, philosophersCount)].BeingUsed = true
			chopsticks[helper.MakeIndex(val+1, philosophersCount)].BeingUsed = true

			fmt.Println("Started Eating:", val)
			time.Sleep(time.Millisecond * helper.GenerateRandTime(50, 500)) // random number for variable finish eating time
			ch2 <- val
		}
	}

}

func think(ch2 <-chan int, ch1 chan<- int) {

	for {
		val, ok := <-ch2
		if !ok {
			fmt.Println("Message delivery Bug. Exiting.")
			return
		}

		chopsticks[helper.MakeIndex(val, philosophersCount)].BeingUsed = false
		chopsticks[helper.MakeIndex(val+1, philosophersCount)].BeingUsed = false

		selected := helper.SelectEater(val, philosophersCount, chopsticks, philosophers)

		if selected >= 0 {
			ch1 <- selected
			println(helper.JoinStrings("philosopher:", strconv.Itoa(val), "Has Finished Eating. Now:", strconv.Itoa(selected), " Can eat"))
		} else {
			println(helper.JoinStrings("philosopher:", strconv.Itoa(val), "Has Finished Eating. But No neighbor can start"))
		}

	}
}
