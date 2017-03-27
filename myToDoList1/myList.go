package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

//TaskCondition defines whether the task is completed or not
type TaskCondition struct {
	Tasks map[string]ToDoList
}

//ToDoList contains the daily tasks
type ToDoList struct {
	Task []string `json:"task"`
	Done bool     `json:"done"`
}

func main() {
r:= mux.NewRouter()
r.HandleFunc("/users/{todotask}", func(w http.ResponseWriter, r *http.Request) { 

	vars := mux.Vars(r)
	if taskinfo, ok := vars["todotask"] ; ok { 
    Tasks := GetTasks("./tasks.json")
	if Tasks.Tasks == nil{
		Tasks = &TaskCondition{map[string]ToDoList{}}
	}


		
	}
}

}

func GetTasks(path string) *TaskCondition{
tasks := &TaskCondition{}
file, e := ioutil.ReadFile(path)
if e != nil { 
	panic(fmt.Sprint("no tasks.json :", e))
}
json.Unmarshal(file, tasks)
return tasks

}




