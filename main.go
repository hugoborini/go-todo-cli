package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

	
type Todo struct {
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

type Todos []Todo


func jSONToTab(pathToJSON string) Todos{
	file, err := ioutil.ReadFile(pathToJSON)
    if err != nil {
      fmt.Print(err)
	}
	
	sb := string(file)

	data := Todos{}
	_ = json.Unmarshal([]byte(sb), &data)

	return data

}

func listTodo(){
	data := jSONToTab("data/todo.json")
	var isOk string
	
	for i := 0; i < len(data); i++{
		if data[i].Done == true {
			isOk = "✅"
		} else{
			isOk = "❌"
		}
		fmt.Println("\nTodo : ", data[i].Todo, ":", isOk, "\n")
	}

}

func addTodo(action string){
	data := jSONToTab("data/todo.json")
	newStruct := &Todo{
		Todo: action,
	}

	data = append(data, *newStruct)

	dataBytes, err := json.MarshalIndent(data, "", "    ")


	err = ioutil.WriteFile("data/todo.json", dataBytes, 0644)
	if err != nil {
        fmt.Println("eroor")
    }
}

func checkTodo(action string){
	data := jSONToTab("data/todo.json");

	var index int

	for i:= 0; i < len(data); i++{
		if(action == data[i].Todo){
			index = i
			
		}
	}

	data[index].Done = true

	dataBytes, err := json.MarshalIndent(data, "", "    ")

	err = ioutil.WriteFile("data/todo.json", dataBytes, 0644)
	if err != nil {
        fmt.Println("error")
    }
}

func deleteTodo(action string){
	data := jSONToTab("data/todo.json")


	for i:= 0; i < len(data); i++{
		if(action == data[i].Todo){
			data= append(data[:i], data[i+1:len(data)]...)

		}
	}
	
	dataBytes, err := json.MarshalIndent(data, "", "    ")
	err = ioutil.WriteFile("data/todo.json", dataBytes, 0644)
	if err != nil {
        fmt.Println("error")
    }

}

func main()  {
	//checkTodo("test 3")

	//addTodo("test 3")
	//listTodo()
	//deleteTodo("test 3")
}