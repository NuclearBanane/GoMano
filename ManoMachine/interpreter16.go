package ManoMachine

import (
	"fmt"
)

func start() (err Error) {
	bus = make(chan int16)
	bus <- 0
	ac = make(chan int16)
	ac <- 0
	dr = make(chan int16)
	dr <- 0
	ar = make(chan int16)
	ar <- 0
	tr = make(chan int16)
	tr <- 0
	mem = make(chan int16)
	mem <- 0
	inr = make(chan int16)
	inr <- 0
	outr = make(chan int16)
	outr <- 0
	pc = make(chan int16)
	pc <- 0
	alu = make(chan int16)
}
