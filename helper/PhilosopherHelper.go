package helper

import "github.com/virtouso/DiningPhilosophersGolang-/models"

func SelectEater(currentPhilosopher int, philosophersCount int, chopsticks []*models.Chopstick, philosophers []*models.Philosopher) int {

	if chopsticks[MakeIndex(currentPhilosopher-1, philosophersCount)].BeingUsed &&
		chopsticks[MakeIndex(currentPhilosopher+2, philosophersCount)].BeingUsed {
		return -1
	}

	if chopsticks[MakeIndex(currentPhilosopher-1, philosophersCount)].BeingUsed {
		return MakeIndex(currentPhilosopher+1, philosophersCount)
	}

	if chopsticks[MakeIndex(currentPhilosopher+2, philosophersCount)].BeingUsed {
		return MakeIndex(currentPhilosopher-1, philosophersCount)
	}

	if philosophers[MakeIndex(currentPhilosopher+1, philosophersCount)].EatCount >= philosophers[MakeIndex(currentPhilosopher-1, philosophersCount)].EatCount {
		return MakeIndex(currentPhilosopher-1, philosophersCount)
	}

	return MakeIndex(currentPhilosopher+1, philosophersCount)
}
