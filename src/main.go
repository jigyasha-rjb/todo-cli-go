package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct{
	Description	string `json:"description"`
	Completed	bool `json:"completed"`
}

const taskFile = "task.json"

func main(){
	//Infinite for loop
	for {
		fmt.Println("TASK LIST")
		fmt.Println("1. Add a Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Mark Task as Complete")
		fmt.Println("5. Exit the application")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')

		choice = strings.TrimSpace(choice)

		switch choice{
			case "1":
				addTask()
			case "2":
				viewTask()
			case "3":
				deleteTask()
			case "4":
				markTaskAsComplete()
			case "5":
				fmt.Println("Exiting ...")
			default:
				fmt.Println("Invalid Choice. Enter number from 1-5.")
				fmt.Println("Please Try Again!")

		}
	}
}


func addTask(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Task Description: ")
	description, _ := reader. ReadString('\n')
	description = strings.TrimSpace(description)
}


	