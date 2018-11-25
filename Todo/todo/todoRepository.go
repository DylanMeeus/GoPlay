package todo

import (
    "encoding/json"
    "io/ioutil"
)

type TodoRepository interface{
    List() *[]Todo
    Save(Todo) error
}

type JsonRepo struct {
    File string
}

func (jr JsonRepo) getTodos() []Todo {
    var todos []Todo
    data, err := ioutil.ReadFile(jr.File)
    if err != nil {
        return make([]Todo, 0)
    }
    err = json.Unmarshal(data, &todos)
    if err != nil {
        panic(err)
    }
    return todos
}

func (jr JsonRepo) List() *[]Todo {
    return nil
}

func (jr JsonRepo) Save(t Todo) error {
    todos := append(jr.getTodos(), t)
    jsonBytes, err := json.Marshal(todos)
    if err != nil {
        return err
    }
    ioutil.WriteFile(jr.File, jsonBytes, 0644)
    return nil
}
