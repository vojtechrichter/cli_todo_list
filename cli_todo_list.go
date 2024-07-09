package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
)

func PromptAddTask() Task {
    reader := bufio.NewReader(os.Stdin) 
    fmt.Printf("Add task\n\n")
    fmt.Printf("Enter task ID: ")

    id, err := reader.ReadBytes('\n')

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\n")

    fmt.Printf("Enter task title: ")

    title, err := reader.ReadBytes('\n')

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\n")

    fmt.Printf("Enter task description: ")

    description, err := reader.ReadBytes('\n')

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\n")

    fmt.Printf("Enter due date: ")

    dueDate, err := reader.ReadBytes('\n')

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\n")

    fmt.Printf("Enter task priority: ")

    priority, err := reader.ReadBytes('\n')

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\n")

    return Task{
        Completed: false,
        Title: string(title)[:len(title)-1],
        Description: string(description)[:len(description)-1],
        DueDate: string(dueDate)[:len(dueDate)-1],
        Priority: string(priority)[:len(priority)-1],
        ID: string(id)[:len(id)-1],
    }
}

func PromptChangeCompletionStatus(complete bool) ([]byte, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("\n")

    if complete {
        fmt.Printf("Enter the id of the task you want to complete: ")
    } else {
        fmt.Printf("Enter the id of the task you want to uncomplete: ")
    }

    id, err := reader.ReadBytes('\n')
    if err != nil {
        return nil, err
    }

    return id[:len(id)-1], nil
}

func PromptRemoveTask() ([]byte, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("\n")

    fmt.Printf("Enter the id of the task you want to remove from the list: ")

    id, err := reader.ReadBytes('\n')
    if err != nil {
        return nil, err
    }

    return id[:len(id)-1], nil
}

func PromptUsage() {
    fmt.Printf("%s show|complete|uncomplete|add|remove\n", os.Args[0]) 
}

func main() {
    tl, err := GetTaskListFromJson(TASKS_FILE)
    if err != nil {
        log.Fatal(err)
    }

    if len(os.Args) < 2 {
        PromptUsage()
        os.Exit(1)
    }

    for _, v := range os.Args {
        switch v {
        case "add": {
            tl.AddTask(PromptAddTask())
        }
        case "complete": {
            id, err := PromptChangeCompletionStatus(true)
            if err != nil {
                log.Fatal(err)
            }
            success := tl.UpdateTaskCompletion(string(id), true)
            fmt.Println(success)
        }
        case "uncomplete": {
            id, err := PromptChangeCompletionStatus(false)
            if err != nil {
                log.Fatal(err)
            }
            _ = tl.UpdateTaskCompletion(string(id), false)
        }
        case "remove": {
            id, err := PromptRemoveTask()
            if err != nil {
                log.Fatal(err)
            }
            _ = tl.RemoveTask(string(id))
        }
        case "show": {
            tl.Display()
        }
        }
    }

    err = SaveTaskList(tl)
    if err != nil {
        log.Fatal(err) 
    }
}
