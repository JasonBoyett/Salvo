// This is meant to demo the current state of the project.
// Feel free to change or delete this when the CLI application is ready.
// If you do modify or delete this file, please update the README.md
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	runner "github.com/JasonBoyett/salvo/local/runner"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nEnter the url to test:")
	address, _ := reader.ReadString('\n')
  address = strings.Trim(address, "\n")

	fmt.Println("Enter the number of users to create:")
	countInput, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.Trim(countInput, "\n"))
	if err != nil {
		fmt.Println("Invalid input.")
		panic(err)
	}

	fmt.Println("\nEnter the number of rquests each user should make per second:")
	rateInput, _ := reader.ReadString('\n')
	rate, err := strconv.ParseFloat(strings.Trim(rateInput, "\n"), 64)
	if err != nil {
		fmt.Println("Invalid input.")
		panic(err)
	}

	fmt.Println("\nEnter the duration of the test in seconds:")
	durationInput, _ := reader.ReadString('\n')
	durationInt, err := strconv.Atoi(strings.Trim(durationInput, "\n"))
  duration := time.Duration(durationInt) * time.Second
	if err != nil {
		fmt.Println("Invalid input.")
		panic(err)
	}

	fmt.Println("\nEnter the number of seconds before requests should timeout:")
	timeoutInput, _ := reader.ReadString('\n')
	timeout, err := strconv.Atoi(strings.Trim(timeoutInput, "\n"))
	if err != nil {
		fmt.Println("Invalid input.")
		panic(err)
	}

	opts := runner.Opts{
		Users:        count,
		Path:         address,
		Rate:         rate,
		Timeout:      timeout,
		Time:         duration,
		SuccessCodes: []int{200},
	}

	results, fails := runner.Run(opts)

	total := len(results)
	fmt.Printf("\nTotal requests: %d\n", total)
	fmt.Printf("Successful requests: %d\n", total-fails)
	fmt.Printf("Failed requests: %d\n", fails)

}
