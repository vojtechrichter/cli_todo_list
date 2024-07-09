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


func main() {
    tl := new(TaskList)

    tl.AddTask(Task{
        Completed: false,
        Title: "Testovaci task",
        Description: "Tohle je examle popis ukolu",
        DueDate: "2024-07-09:10-00-00",
        Priority: TASK_PRIORITY_CRITICAL,
        ID: "A32",
    })

    tl.AddTask(Task{
        Completed: true,
        Title: "Implement content caching",
        Description: "Set up redis for client-side caching",
        DueDate: "2024-07-21:not_specified",
        Priority: TASK_PRIORITY_NORMAL,
        ID: "REDIS234",
    })

    tl.AddTask(PromptAddTask())

    err := SaveTaskList(tl)
    if err != nil {
        log.Fatal(err) 
    }

    tl.Display()
}
