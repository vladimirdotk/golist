package main

import (
	"fmt"
	"sync"
)

// To get a feel for what it’s like to use Broadcast,
// let’s imagine we’re creating a GUI application with a button on it.
// We want to register an arbitrary number of functions that will run when that button is clicked.
// A Cond is perfect for this because we can use its Broadcast method to notify all registered handlers.

func main()  {
	type Button struct { Clicked *sync.Cond}

	button := Button{ Clicked: sync.NewCond(&sync.Mutex{}) }
	// Here we define a convenience function
	// that will allow us to register functions to handle signals from a condition.
	// Each handler is run on its own goroutine,
	// and subscribe will not exit until that goroutine is confirmed to be running.
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})
	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
