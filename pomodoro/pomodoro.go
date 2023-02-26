package main

import (
	"bufio"
	"fmt"
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
		return 0, fmt.Errorf("invalid input, please enter a number")
	}
	return time.Duration(inputInt) * time.Minute, nil
}

// Prompt the user to enter loops
func promptLoops(label string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the number of times you want to %v the pomodoro: ", label)
	inputStr, _ := reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)
	inputInt, err := strconv.Atoi(inputStr)
	if err != nil {
		return 0, fmt.Errorf("hat is an invalid numer, please enter a valid number")
	}
	return int(inputInt), nil
}

func runThepomodoro(loops int, pomodoro *Pomodoro) {
	// Adding a progress bar
	bar := progressbar.NewOptions(
		100,
		progressbar.OptionSetTheme(progressbar.Theme{Saucer: "📝", SaucerPadding: "", BarStart: "🏁", BarEnd: "🏁"}),
		progressbar.OptionShowBytes(true),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetRenderBlankState(true),
	)
	for i := 0; i < loops; i++ {
		// Start the Pomodoro timer
		pomodoro.Start()

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

		// Update the progress bar for the break
		bar.ChangeMax64(int64(pomodoro.BreakDuration.Seconds()))
		// bar.Set64(int64(pomodoro.BreakDuration.Seconds()))
		time.Sleep(pomodoro.BreakDuration)
	}
}

func main() {
	// Introducing the pomodoro app in the cli
	fmt.Println("Welcome to the Pomodoro app!")

	// Prompt the user to enter the duration and break duration of the Pomodoro timer
	duration, err := promptDuration("Pomodoro duration")
	if err != nil {
		fmt.Println(err)
		return
	}
	breakDuration, err := promptDuration("break duration")
	if err != nil {
		fmt.Println(err)
		return
	}

	loops, err := promptLoops("loop through")
	if err != nil {
		fmt.Println(err)
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

	runThepomodoro(loops, &pomodoro)

	fmt.Printf("The pomodoro app did %v loops. I hope you studied well!\n", loops)
	fmt.Println("Goodbye!")
	os.Exit(0)

}
