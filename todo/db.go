package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Todo struct {
	ID        int    
	Text      string `json:"text"`
	IsDone    bool   `json:"is_done"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var todos []Todo
var currentID int = 0

func StoreTodo(task string, filepath string) {
	// Load existing todos from file
	data, err := os.ReadFile(filepath)
	if err == nil {
		_ = json.Unmarshal(data, &todos)
		for _, todo := range todos {
			if todo.ID > currentID {
				currentID = todo.ID
			}
		}
	} else if !os.IsNotExist(err) {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Create new todo
	currentID++
	newTodo := Todo{
		ID:        currentID,
		Text:      task,
		IsDone:    false,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	// Append and store
	todos = append(todos, newTodo)
	response, err := json.MarshalIndent(todos, "", "\t")
	if err != nil {
		fmt.Printf("Error marshaling todos: %v\n", err)
		return
	}

	err = os.WriteFile(filepath, response, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("Todo item stored successfully!")
}


func GetTodos(filename string) []Todo {

	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Todo file not found. Returning empty list.")
			return []Todo{}
		}
		fmt.Printf("Error reading file: %v\n", err)
		return []Todo{}
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return []Todo{}
	}

	// Display todos using tabwriter
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tTASK\tSTATUS\tCREATED_AT\tUPDATED_AT")
	for _, todo := range todos {
		
		status := "Pending"
		if todo.IsDone {
			status = "Done"
		}
		fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\n", todo.ID, todo.Text, status, todo.CreatedAt, todo.UpdatedAt)
	}
	writer.Flush()

	return todos
}
func MarkTodoAsDone(id int, filepath string) {
	updated := false

	// Load todos from file
	data, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Todo file not found. Returning empty list.")
			return
		}
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Mark the todo as done
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].IsDone = true
			todos[i].UpdatedAt = time.Now().Format(time.RFC3339)
			updated = true
			break
		}
	}

	if !updated {
		fmt.Printf("Todo with ID %d not found.\n", id)
		return
	}

	// Save updated todos back to file
	data, err = json.MarshalIndent(todos, "", "\t")
	if err != nil {
		fmt.Printf("Error marshaling updated todos: %v\n", err)
		return
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Printf("Error writing updated todos to file: %v\n", err)
		return
	}

	fmt.Printf("Todo #%d marked as done successfully!\n", id)
}

func DeleteTodo(id int, filepath string) {
	
	updated := false
//at first we need to load the todos from the file
	data, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Todo file not found. Returning empty list.")
			return
		}
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			fmt.Printf(("todo id %d and given id %d\n"), todo.ID, id)
			todos = append(todos[:i], todos[i+1:]...)
			updated = true
			break
		}
	}

	if !updated {
		fmt.Printf("Todo with ID %d not found.\n", id)
		return
	}

	data, err = json.MarshalIndent(todos, "", "\t")
	if err != nil {
		fmt.Printf("Error marshaling updated todos: %v\n", err)
		return
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Printf("Error writing updated todos to file: %v\n", err)
		return
	}

	fmt.Printf("Todo #%d deleted successfully!\n", id)
}
