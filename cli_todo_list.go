package main

import (
    "fmt"
    _ "os"
)

const (
    TASK_PRIORITY_NORMAL = "normal"
    TASK_PRIORITY_HIGH = "high"
    TASK_PRIORITY_MAJOR = "major"
    TASK_PRIORITY_CRITICAL = "critical"
)

type Task struct {
    Completed bool
    Title string
    Description string
    DueDate string
    Priority string
    ID string
}

func (t Task) Display() {
    if t.Completed {
        fmt.Printf("[%s] %s (%s)\n", "x", t.Title, t.ID)
    } else {
        fmt.Printf("[%s] %s (%s)\n", "-", t.Title, t.ID)
    }
    fmt.Printf("\t%s\n\n", t.Description)
    fmt.Printf("\tDue date: %s\n", t.DueDate)
    fmt.Printf("\tPriority: %s\n", t.Priority)
}

type TaskList []Task

func (tl *TaskList) AddTask(t Task) {
    *tl = append(*tl, t)
}

func (tl *TaskList) Display() {
    fmt.Printf("\n")
    for _, t := range *tl {
        t.Display()
    }
    fmt.Printf("\n")
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

    tl.Display()
}
