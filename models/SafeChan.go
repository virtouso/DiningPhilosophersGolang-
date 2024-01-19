package models

type SafeChan struct {
	Ch     chan int
	IsOpen bool
}
