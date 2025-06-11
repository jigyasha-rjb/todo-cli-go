package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct{
	Description	string `json:"description"`
	Completed	bool `json:"completed"`
}

const taskFile = "task.json"

func main(){
	fmt.Println("----------------------")
	fmt.Println("CLI TO-DO APLLICATION")
	fmt.Println("----------------------")
	//Infinite for loop
	for {
		fmt.Println("\nOPTIONS")
		fmt.Println("1. Add a Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Mark Task as Complete")
		fmt.Println("5. Exit the application")
		fmt.Println("------------------------")
		fmt.Println("Enter your choice: ")

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
				return
			default:
				fmt.Println("Invalid Choice. Enter number from 1-5.")
				fmt.Println("Please Try Again!")
				fmt.Println("------------------------")
		}
	}
}

func loadTasks()([]Task, error){
	// If no contents in taskFile
	if _, err := os.Stat(taskFile); os.IsNotExist(err){
		return []Task{}, nil
		//return []Task{}, nil
	}
	
	//If there is content is task file
	data, err := os.ReadFile(taskFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read task file: %w", err)
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil{
		return nil, fmt.Errorf("failed to unmarshal task from JSON file: %w", err)
	}
	return tasks, nil

}

func saveTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON file: %w", err)
	}
	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks to file: %w", err)
	}
	return nil
}

func addTask(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------------------")
	fmt.Println("Enter Task Description: ")
	description, _ := reader. ReadString('\n')
	description = strings.TrimSpace(description)

	// If error is encountered
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}
	tasks = append(tasks, Task{Description: description, Completed: false})
	
	err = saveTasks(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}
	fmt.Println("\nTASK ADDED!")
	fmt.Println("----------------------")
}

func viewTask(){
	tasks, err := loadTasks()
	if err != nil { 
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("-------------------------------------")
	fmt.Printf("%-5s %-20s %s\n","ID", "Description","Completed")
	fmt.Println("-------------------------------------")
	for i, task:= range tasks{
		fmt.Printf("%5d %-20s %t\n", i+1, task.Description, task.Completed)
	}
	fmt.Println("-------------------------------------")
}

func deleteTask(){
	tasks, err := loadTasks()
	if err != nil { 
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	viewTask()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the task to delete: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskNum, err := strconv.Atoi(input)

	if err != nil || taskNum < 1 || taskNum > len(tasks){
		fmt.Println("Invalid Task ID")
	}

	tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)

	err = saveTasks(tasks)
	if err != nil{
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}
	fmt.Println("\nTASK DELETED!")
	fmt.Println("----------------------")
}

func markTaskAsComplete(){
	tasks, err := loadTasks()
	if err != nil { 
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	viewTask()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the task to mark as completed: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskNum, err := strconv.Atoi(input)

	if err != nil || taskNum < 1 || taskNum > len(tasks){
		fmt.Println("Invalid Task ID")
	}

	tasks[taskNum-1].Completed = true

	err = saveTasks(tasks)
	if err != nil{
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}
	fmt.Println("\nTASK MARKED AS COMPLETED!")
	fmt.Println("----------------------------")
}