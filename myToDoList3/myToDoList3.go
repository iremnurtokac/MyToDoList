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

	r.HandleFunc("/users/{task}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		if taskName, ok := vars["task"]; ok {

			taskList := GetTasks("./tasks.json")

			if taskList.Tasks == nil {
				taskList = &TasksToDo{map[string]MyToDoList{}}
			}
			td, ok := taskList.Tasks[taskName]
			if !ok {
				td = MyToDoList{[]string{}, ""}

				taskList.Tasks[taskName] = td

			}
			taskList.Tasks[taskName] = td
			WriteTasks("./tasks.json", taskList)
			if currentAsJSON, err := json.Marshal(td); err == nil {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(http.StatusOK)
				w.Write(currentAsJSON)
				//fmt.Fprintf(w, "%v", string(currentAsJSON))

			} else {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error while creating JSON : %v", err)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Please provide Game in URL")
		}

	}).Methods("GET")

	r.HandleFunc("/users/{task}/move/{task:[1-9]}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		if taskName, ok := vars["task"]; ok {

			taskList := GetTasks("./tasks.json")

			if taskList.Tasks == nil {
				taskList = &TasksToDo{map[string]MyToDoList{}}
			}
			td, ok := taskList.Tasks[taskName]
			if !ok {
				td = MyToDoList{[]string{}, ""}

				taskList.Tasks[taskName] = td

			}

			taskList.Tasks[taskName] = td
			WriteTasks("./tasks.json", taskList)
			if currentAsJSON, err := json.Marshal(td); err == nil {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(http.StatusOK)
				w.Write(currentAsJSON)
				//fmt.Fprintf(w, "%v", string(currentAsJSON))

			} else {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error while creating JSON : %v", err)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Please provide Game in URL")
		}

	}).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func (td *MyToDoList) addingTask(taskNum int) {
	taskNum = 0
	for taskNum = range td.Task {

		if len(td.Task[taskNum]) < 1 {
			td.Task[taskNum] = td.Info
		} else {
			td.Info = td.Task[taskNum]

		}
		taskNum++

	}

}

func GetTasks(path string) *TasksToDo {

	tasks := &TasksToDo{}
	list, e := ioutil.ReadFile(path)
	if e != nil {
		panic(fmt.Sprint("no tasks.json :", e))
	}
	json.Unmarshal(list, tasks)

	return tasks
}

func WriteTasks(path string, taskList *TasksToDo) {

	d1, _ := json.Marshal(*taskList)
	ioutil.WriteFile(path, d1, 0755)

}
