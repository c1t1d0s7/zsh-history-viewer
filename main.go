package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	historyFilePath := os.Getenv("HOME") + "/.zsh_history"

	readFile, err := os.Open(historyFilePath)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	// Move to the end of the file
	_, err = readFile.Seek(0, io.SeekEnd)
	if err != nil {
		log.Fatal(err)
	}

	// Get the current time
	now := time.Now()

	// Format the current time as YYMMDD_HHmm
	formattedTimeNow := now.Format("060102_1504")

	// Create the file name
	outputFileName := fmt.Sprintf("history_%s.txt", formattedTimeNow)

	// Create a new file
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	history := bufio.NewReader(readFile)
	for {
		line, err := history.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			time.Sleep(100 * time.Millisecond) // wait for 0.1 second
			continue
		}

		// Split the line into timestamp, duration, and command
		s := strings.Split(line, ";")

		if len(s) != 2 {
			fmt.Println("Invalid line format")
			continue
		}

		timestampAndDuration := s[0]
		p := strings.Split(timestampAndDuration, ":")

		if len(p) < 3 {
			fmt.Println("Invalid timestamp and duration format")
			continue
		}

		timestampStr := strings.TrimSpace(p[1])
		// duration := p[2]
		command := s[1]

		// Convert the timestamp to a human-readable format
		timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			fmt.Printf("Failed to parse timestamp: %v\n", err)
			continue
		}
		timestampAsTime := time.Unix(timestampInt, 0)
		formattedTime := timestampAsTime.Format("15:04:05")

		// fmt.Printf("Timestamp: %s, Duration: %s, Command: %s\n", timestampAsTime.Format(time.RFC3339), duration, command)
		fmt.Printf("[%s] %s", formattedTime, command)
		fmt.Fprintf(outputFile, "[%s] %s", formattedTime, command)
	}
}
