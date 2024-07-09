package main

import (
    "fmt"
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
