package main

import (
    "fmt"
    "slices"
)

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
