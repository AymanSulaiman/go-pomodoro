package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

// Define the Pomodoro struct
type Pomodoro struct {
	Duration      time.Duration // The duration of the Pomodoro timer
	BreakDuration time.Duration // The duration of the break after each Pomodoro
	Running       bool          // Whether the Pomodoro timer is currently running or not
	StartTime     time.Time     // The time when the Pomodoro timer started
	EndTime       time.Time     // The time when the Pomodoro timer ends
}

// Start the Pomodoro timer
func (p *Pomodoro) Start() {
	p.Running = true
	p.StartTime = time.Now()
	p.EndTime = p.StartTime.Add(p.Duration)
	fmt.Printf("Pomodoro timer started for %v.\n", p.Duration)
}

// End the Pomodoro timer
func (p *Pomodoro) End() {
	p.Running = false
	p.StartTime = time.Time{}
	p.EndTime = time.Time{}
	fmt.Println("Pomodoro timer ended.")
}

// Prompt the user to enter a duration
func promptDuration(label string) (time.Duration, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s (in minutes): ", label)
	inputStr, _ := reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)
	inputInt, err := strconv.Atoi(inputStr)
	if err != nil {
		return 0, fmt.Errorf("that is an invalid number, please enter a number")
	}
	if inputInt <= 0 {
		return 0, fmt.Errorf("that is an invalid number, please enter a positive integer")
	}
	return time.Duration(inputInt) * time.Minute, nil
}

// Prompt the user to enter loops
func promptLoops() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of times you want to loop through the pomodoro: ")
	inputStr, _ := reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)
	inputInt, err := strconv.Atoi(inputStr)
	if err != nil {
		return 0, fmt.Errorf("that is an invalid number, please enter a valid number")
	}

	if inputInt <= 0 {
		return 0, fmt.Errorf("that is an invalid number, please enter a positive integer")
	}
	return int(inputInt), nil
}

// Core Pomodoro Function
func runThePomodoro(loops int, pomodoro *Pomodoro) {
	for i := 0; i < loops; i++ {
		fmt.Printf("\nThis is study round: %v\n", i+1)
		// Start the Pomodoro timer
		pomodoro.Start()

		// Adding a progress bar
		bar := progressbar.DefaultBytes(
			int64(pomodoro.Duration.Seconds()),
			"pomodoro progress",
		)

		// Wait for the Pomodoro timer to end
		for pomodoro.Running {
			// Update the progress bar
			elapsed := time.Since(pomodoro.StartTime).Seconds()
			if elapsed > pomodoro.Duration.Seconds() {
				elapsed = pomodoro.Duration.Seconds()
			}
			bar.Set64(int64(elapsed))

			if time.Now().After(pomodoro.EndTime) {
				pomodoro.End()
			}
			time.Sleep(1 * time.Second)
		}

		// Start the break timer
		fmt.Printf("Break started for %v.\n", pomodoro.BreakDuration)

		// Reset the progress bar for the break
		bar.Reset()

		// Adding a progress bar for the break
		breakBar := progressbar.DefaultBytes(
			int64(pomodoro.BreakDuration.Seconds()),
			"break progress",
		)

		// Update the progress bar for the break
		breakBar.Set64(int64(pomodoro.BreakDuration.Seconds()))
	}
}

func doTheWholePomodoro() {
	// Introducing the pomodoro app in the cli
	fmt.Println("Welcome to the Pomodoro app!")

	// Prompt the user to enter the duration and break duration of the Pomodoro timer
	duration, err := promptDuration("Pomodoro duration")
	if err != nil {
		log.Println(err)
		return
	}
	breakDuration, err := promptDuration("Break duration")
	if err != nil {
		log.Println(err)
		return
	}

	loops, err := promptLoops()
	if err != nil {
		log.Println(err)
		return
	}

	// Create a new Pomodoro timer
	pomodoro := Pomodoro{
		Duration:      duration,
		BreakDuration: breakDuration,
		Running:       false,
		StartTime:     time.Time{},
		EndTime:       time.Time{},
	}

	runThePomodoro(loops, &pomodoro)

	fmt.Printf("\nThe pomodoro app did %v loops. I hope you studied well!\n", loops)
}

func main() {
	// Starts the pomodoro timer and prompts the user to start another pomodoro, with input validation for "y" or "n".
	for {
		doTheWholePomodoro()

		fmt.Print("Do you want to do the pomodoro again? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.ToLower(strings.TrimSpace(answer))
		if answer == "y" {
			fmt.Println("Continuing...")
			break
		} else if answer == "n" {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}

	fmt.Println("Goodbye!")
	os.Exit(0)

}
