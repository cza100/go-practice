package main

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Println(balance)
		case balances <- balance:
			fmt.Println("rrr")
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {
	go func() {
		Deposit(200)                 // A1
		fmt.Println("=1", Balance()) // A2
	}()

	// Bob:
	go Deposit(100)
	fmt.Println("=2", Balance())
}
