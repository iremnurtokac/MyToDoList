package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//TasksToDo is the list of tasks
type TasksToDo struct {
	Tasks map[string]MyToDoList
}

//MyToDoList defines tasks to do and whether they are done or not
type MyToDoList struct {
	Task []string `json:"task"`
	Info string   `json:"info"`
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {

		taskList := GetTasks("./tasks.json")

		if taskList == nil || len(taskList.Task) == 0 {
			taskList = &MyToDoList{[]string{}, ""}
		}

		if currentAsJSON, err := json.Marshal(taskList); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			w.Write(currentAsJSON)
			//fmt.Fprintf(w, "%v", string(currentAsJSON))

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error while creating JSON : %v", err)
		}

	}).Methods("GET")

	r.HandleFunc("/tasks/add", func(w http.ResponseWriter, r *http.Request) {

		taskList := GetTasks("./tasks.json")
		r.ParseForm()
		taskList.Task = append(taskList.Task, r.FormValue("task"))

		fmt.Println(r.FormValue("task"))
		fmt.Println(taskList.Task)
		taskList.addingInfo()

		WriteTasks("./tasks.json", taskList)
		if currentAsJSON, err := json.Marshal(taskList); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			w.Write(currentAsJSON)
			//fmt.Fprintf(w, "%v", string(currentAsJSON))

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error while creating JSON : %v", err)
		}

	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func (td *MyToDoList) addingInfo() {
	task := td.Info
	taskNum := 0
	for taskNum = range td.Task {

		if len(td.Task[taskNum]) < 1 {
			td.Task[taskNum] = task

		}
		taskNum++

	}
}

// func (td *MyToDoList) addingTask() {
// 	taskNum := 0
// 	for taskNum = range td.Task {

// 		if len(td.Task[taskNum]) < 1 {
// 			td.Task[taskNum] = td.Info

// 		}
// 		taskNum++

// 	}

// }

func GetTasks(path string) *MyToDoList {

	tasks := &MyToDoList{}
	list, e := ioutil.ReadFile(path)
	if e != nil {
		panic(fmt.Sprint("no tasks.json :", e))
	}
	json.Unmarshal(list, tasks)

	return tasks
}

func WriteTasks(path string, taskList *MyToDoList) {

	d1, _ := json.Marshal(*taskList)
	ioutil.WriteFile(path, d1, 0755)

}
