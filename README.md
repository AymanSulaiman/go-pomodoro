# go-pomodoro

This is a Pomodoro Timer Command Line Interface (CLI) app implemented in Go. It allows users to set a study time and a break time, and loop through them a specified number of times. The app uses the schollz/progressbar package to display a progress bar during the Pomodoro timer and break.

## How to use

1. Download and install Go from the official website.
    * macOs: `brew install go`
    * Linux:  `rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz` then ` rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz` then to verify `go version`
    * Windows: Just go here -> `https://go.dev/doc/install`

2. Clone the repository: `git clone git@github.com:AymanSulaiman/go-pomodoro.git`

3. Navigate to the cloned directory: `cd go-pomodoro`

4. Run the app: `go run main.go`

5. The app will prompt you to enter the duration of the Pomodoro timer and the duration of the break.

6. The app will also prompt you to enter the number of times you want to loop through the Pomodoro timer and break.

7. The app will display a progress bar during the Pomodoro timer and break.

8. After each loop, the app will prompt you to continue or exit.

## Dependencies

`github.com/schollz/progressbar/v3`