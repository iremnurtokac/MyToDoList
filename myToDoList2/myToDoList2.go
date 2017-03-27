package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//TasksToDo is the list of tasks
type TasksToDo struct {
	Tasks map[string]MyToDoList
}

//MyToDoList defines tasks to do and whether they are done or not
type MyToDoList struct {
	Task [9]string `json:"task"`
	Info string    `json:"info"`
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/users/{task}", func(w http.ResponseWriter, r *http.Request) {
		/*

		   GET http://localhost:8080/users/Tictactoe HTTP/1.1

		*/
		vars := mux.Vars(r)
		if taskName, ok := vars["task"]; ok {

			taskList := GetTasks("./tasks.json")

			if taskList.Tasks == nil {
				taskList = &TasksToDo{map[string]MyToDoList{}}
			}
			ts, ok := taskList.Tasks[taskName]
			if !ok {
				ts = MyToDoList{[9]string{}, ""}

				taskList.Tasks[taskName] = ts

			}
			taskList.Tasks[taskName] = ts
			WriteTasks("./tasks.json", taskList)
			if currentAsJSON, err := json.Marshal(ts); err == nil {
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

	r.HandleFunc("/users/{task}/move/{field:[1-9]}", func(w http.ResponseWriter, r *http.Request) {
		/*

		   POST http://localhost:8080/users/Tictactoe/move/3 HTTP/1.1
		   POST http://localhost:8080/users/Tictactoe/move/4 HTTP/1.1
		   POST http://localhost:8080/users/Tictactoe/move/2 HTTP/1.1

		*/

		vars := mux.Vars(r)
		if taskName, ok := vars["task"]; ok {

			taskList := GetTasks("./tasks.json")

			if taskList.Tasks == nil {
				taskList = &TasksToDo{map[string]MyToDoList{}}
			}
			ts, ok := taskList.Tasks[taskName]
			if !ok {
				ts = MyToDoList{[9]string{}, ""}

				taskList.Tasks[taskName] = ts

			}

			if ts.addingTasks() {

				if taskstring, ok := vars["task"]; ok {
					if task, err := strconv.Atoi(taskstring); err == nil {
						if ts.Info == "X" {
							ts.Info = "O"
						} else {
							ts.Info = "X"
						}
						ts.playerPut(task - 1)

					}
				}

			}

			ts.addingTasks()
			taskList.Tasks[taskName] = ts
			WriteTasks("./tasks.json", taskList)
			if currentAsJSON, err := json.Marshal(ts); err == nil {
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
	//

	r.HandleFunc("/users/{task}/reset", func(w http.ResponseWriter, r *http.Request) {
		/*

		   DELETE http://localhost:8080/users/reset HTTP/1.1

		*/
		vars := mux.Vars(r)
		if taskName, ok := vars["task"]; ok {

			taskList := GetTasks("./tasks.json")

			ts := MyToDoList{[9]string{}, ""}

			taskList.Tasks[taskName] = ts
			WriteTasks("./tasks.json", taskList)
			if currentAsJSON, err := json.Marshal(ts); err == nil {
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

func (ts *MyToDoList) addingTasks() bool {

	return ts.draw()

}

func (ts *MyToDoList) playerPut(task int) {
	if len(ts.Task[task]) < 1 {
		ts.Task[task] = ts.Info

	} else {

		ts.Info = ts.Task[task]
	}

}

func compare(tocompare ...string) bool {
	res := true
	current := ""
	if len(tocompare) > 0 {
		current = tocompare[0]
	} else {
		return false
	}
	for _, v := range tocompare {
		if len(v) == 0 {
			res = false
			break
		}

		if v != current {
			res = false
			break
		}

	}
	return res

}

func (ts *MyToDoList) draw() bool {
	i := 0
	for _, v := range ts.Task {
		if v == "X" || v == "Y" {
			i++
		}
	}
	if i == 9 {

		//gs.Message = fmt.Sprint("Draw !")

		return true
	}
	return false

}

func GetTasks(path string) *TasksToDo {

	tasks := &TasksToDo{}
	file, e := ioutil.ReadFile(path)
	if e != nil {
		panic(fmt.Sprint("no tasks.json :", e))
	}
	json.Unmarshal(file, tasks)

	return tasks
}

func WriteTasks(path string, taskList *TasksToDo) {

	d1, _ := json.Marshal(*taskList)
	ioutil.WriteFile(path, d1, 0755)

}
