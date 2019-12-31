package main

/*
from:
https://studygolang.com/articles/13866
https://github.com/pagnihotry/golang_samples/blob/master/go_context_sample.go
*/
import (
	"context"
	"fmt"
	"time"
)

func main() {
	//Make a background context
	ctx := context.Background()
	//Derive a context with cancel
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	//defer canceling so that all the resources are freed up
	//For this and the derived contexts
	defer func() {
		fmt.Println("11.Main Defer: canceling context")
		cancelFunction()
	}()
	//Cancel context after a random time
	//This cancels the request after a random timeout
	//If this happens, all the contexts derived from this should return
	go func() {
		sleepRandom("Main", nil)
		cancelFunction()
		fmt.Println("12.Main Sleep complete. canceling context")
	}()
	//Do work
	doWorkContext(ctxWithCancel)
}

//Slow function
func sleepRandom(fromFunction string, ch chan int) {
	//defer cleanup
	defer func() { fmt.Println(fromFunction, "51. sleepRandom complete") }()
	//Perform a slow task
	//For illustration purpose,
	//Sleep here for random ms
	//seed := time.Now().UnixNano()
	//r := rand.New(rand.NewSource(seed))
	//randomNumber := r.Intn(100)
	//sleeptime := randomNumber + 100
	sleeptime := 91 + 100
	fmt.Println(fromFunction, "52. Starting sleep for", sleeptime, "ms")
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "53. Waking up, slept for ", sleeptime, "ms")
	//write on the channel if it was passed in
	if ch != nil {
		ch <- sleeptime
	}
}

//Function that does slow processing with a context
//Note that context is the first argument
func sleepRandomContext(ctx context.Context, ch chan bool) {
	//Cleanup tasks
	//There are no contexts being created here
	//Hence, no canceling needed
	defer func() {
		fmt.Println("31. sleepRandomContext complete")
		ch <- true
	}()
	//Make a channel
	sleeptimeChan := make(chan int)
	//Start slow processing in a goroutine
	//Send a channel for communication
	go sleepRandom("sleepRandomContext", sleeptimeChan)
	//Use a select statement to exit out if context expires
	select {
	case <-ctx.Done():
		//If context is cancelled, this case is selected
		//This can happen if the timeout doWorkContext expires or
		//doWorkContext calls cancelFunction or main calls cancelFunction
		//Free up resources that may no longer be needed because of aborting the work
		//Signal all the goroutines that should stop work (use channels)
		//Usually, you would send something on channel,
		//wait for goroutines to exit and then return
		//Or, use wait groups instead of channels for synchronization
		fmt.Println("32. sleepRandomContext: Time to return")
	case sleeptime := <-sleeptimeChan:
		//This case is selected when processing finishes before the context is cancelled
		fmt.Println("33. Slept for ", sleeptime, "ms")
	}
}

//A helper function, this can, in the real world do various things.
//In this example, it is just calling one function.
//Here, this could have just lived in main
func doWorkContext(ctx context.Context) {
	//Derive a timeout context from context with cancel
	//Timeout in 150 ms
	//All the contexts derived from this will returns in 150 ms
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)
	//Cancel to release resources once the function is complete
	defer func() {
		fmt.Println("21. doWorkContext complete")
		cancelFunction()
	}()
	//Make channel and call context function
	//Can use wait groups as well for this particular case
	//As we do not use the return value sent on channel
	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)
	//Use a select statement to exit out if context expires
	select {
	case <-ctx.Done():
		//This case is selected when the passed in context notifies to stop work
		//In this example, it will be notified when main calls cancelFunction
		fmt.Println("22. doWorkContext: Time to return")
	case <-ch:
		//This case is selected when processing finishes before the context is cancelled
		fmt.Println("23. sleepRandomContext returned")
	}
}
