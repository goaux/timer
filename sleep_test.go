package timer_test

import (
	"context"
	"fmt"
	"time"

	"github.com/goaux/timer"
)

func Example() {
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Try to sleep for 1 second
	fmt.Println("Starting sleep for 1 second")
	err := timer.SleepCause(ctx, 1*time.Second)
	if err != nil {
		fmt.Printf("Sleep interrupted: %v\n", err)
	} else {
		fmt.Println("Sleep completed successfully")
	}

	// Try to sleep for 3 seconds (longer than the context timeout)
	fmt.Println("Starting sleep for 3 seconds")
	err = timer.Sleep(ctx, 3*time.Second)
	if err != nil {
		fmt.Printf("Sleep interrupted: %v\n", err)
	} else {
		fmt.Println("Sleep completed successfully")
	}
	// Output:
	// Starting sleep for 1 second
	// Sleep completed successfully
	// Starting sleep for 3 seconds
	// Sleep interrupted: context deadline exceeded
}

func ExampleSleep() {
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Try to sleep for 1 second
	fmt.Println("Starting sleep for 1 second")
	err := timer.Sleep(ctx, 1*time.Second)
	if err != nil {
		fmt.Printf("Sleep interrupted: %v\n", err)
	} else {
		fmt.Println("Sleep completed successfully")
	}

	// Try to sleep for 3 seconds (longer than the context timeout)
	fmt.Println("Starting sleep for 3 seconds")
	err = timer.Sleep(ctx, 3*time.Second)
	if err != nil {
		fmt.Printf("Sleep interrupted: %v\n", err)
	} else {
		fmt.Println("Sleep completed successfully")
	}
	// Output:
	// Starting sleep for 1 second
	// Sleep completed successfully
	// Starting sleep for 3 seconds
	// Sleep interrupted: context deadline exceeded
}
