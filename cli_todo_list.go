package main

import (
    "fmt"
    "os"
    "slices"
    "bufio"
    "log"
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

func (tl *TaskList) GetTaskById(id string) (Task, bool) {
    for _, v := range *tl {
        if v.ID == id {
            return v, true
        }
    }

    return Task{}, false
}

func (tl *TaskList) UpdateTaskCompletion(id string, complete bool) bool {
    for i, _ := range *tl {
        if (*tl)[i].ID == id {
            (*tl)[i].Completed = complete

            return true
        }
    }

    return false
}

func (tl *TaskList) RemoveTask(id string) bool {
    for i, _ := range *tl {
        if (*tl)[i].ID == id {
            // find a better way
            *tl = slices.Delete(*tl, i, i+1)
            return true
        }
    }

    return false
}

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
        DueDate: string(dueDate)[:len(description)-1],
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

    tl.Display()
}
