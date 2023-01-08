package main

import (
	"fmt"
	"time"
)

func main() {

	tasks := []string{
		"Task 1- rest API",
		"Task 2- go routine",
		"Task 3- PAM",
		"Task 4-Function",
	}
	users := []string{
		"Ashfaque",
		"Prashant",
		"Saheb",
		"Bijay",
	}
	for i, task := range tasks {
		go func(task string, user string) {
			fmt.Printf("Assigning %s to %s\n", task, user)
		}(task, users[i])
	}
	time.Sleep(time.Second)
}
