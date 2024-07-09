package main

import (
    "encoding/json"
    "os"
)

const TASKS_FILE = "./tasks.json"

func GetJsonFromTaskList(tl *TaskList) ([]byte, error) {
    tlJson, err := json.Marshal(tl)
    if err != nil {
        return nil, err
    }

    return tlJson, nil
}

func GetTaskListFromJson(filename string) (*TaskList, error) {
    tl := new(TaskList)

    data, err := os.ReadFile(filename)
    if err != nil {
        return tl, err
    }

    if err = json.Unmarshal([]byte(data), tl); err != nil {
        return tl, err
    }

    return tl, nil
}

func SaveTaskList(tl *TaskList) error {
    json, err := GetJsonFromTaskList(tl)
    if err != nil {
        return err
    }

    f, err := os.Create(TASKS_FILE)
    if err != nil {
        return err
    }

    _, err = f.Write(json)
    if err != nil {
        f.Close()
        return err
    }

    if err = f.Close(); err != nil {
        return err
    }

    return nil
}
