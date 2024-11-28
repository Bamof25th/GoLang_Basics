package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Data struct {
	UserId    int    `json:"user_id"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getRequest() {
	fmt.Println("learning CRUD in go")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/4")

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("response status:", res.Status)

	if res.StatusCode != http.StatusOK {
		fmt.Println("error in getting response:", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Println("Error reading data :", err)
	// 	return
	// }

	// fmt.Println("Data:", string(data))

	var todo Data

	err = json.NewDecoder(res.Body).Decode(&todo)

	if err != nil {
		fmt.Println("Error decoding json:", err)
		return
	}

	fmt.Println("Todo:", todo)
	fmt.Println("Todo:", todo.Title)

}

func postRequest() {
	todo := Data{
		UserId:    1,
		Id:        101,
		Title:     "new task",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling json:", err)
		return
	}
	// covert json data to string
	jsonString := string(jsonData)
	//  convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println(" error while data posting", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("response status:", res.Status)
	if res.StatusCode != http.StatusCreated {
		fmt.Println("error in posting data:", res.Status)
		return
	}
	fmt.Println("Data posted successfully")
	var newTodo Data
	err = json.NewDecoder(res.Body).Decode(&newTodo)
	if err != nil {
		fmt.Println("Error decoding json:", err)
		return
	}
	fmt.Println("New Todo:", newTodo)

}

func putRequest() {
	todo := Data{
		UserId:    1,
		Id:        101,
		Title:     "updated task",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling json:", err)
		return
	}
	// covert json data to string
	jsonString := string(jsonData)
	//  convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/todos/32", jsonReader)
	if err != nil {
		fmt.Println(" error while data updating", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(" error while data updating", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("response status:", res.StatusCode)
	if res.StatusCode != http.StatusOK {
		fmt.Println("error in updating data:", res.Status)
		return
	}
	fmt.Println("Data updated successfully")
	var updatedTodo Data
	err = json.NewDecoder(res.Body).Decode(&updatedTodo)
	if err != nil {
		fmt.Println("Error decoding json:", err)
		return
	}
	fmt.Println("updated todo :", updatedTodo)

}

func deleteRequest() {
	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/todos/32", nil)
	if err != nil {
		fmt.Println("error while data deleting", err)
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error while data deleting", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("response status:", res.StatusCode)
	if res.StatusCode != http.StatusOK {
		fmt.Println("error in deleting data:", res.Status)
		return
	}
	fmt.Println("Data deleted successfully")
}

func main() {
	//  CRUD

	// Get request
	// getRequest();

	//  POST request
	// postRequest()

	// PUT request
	// putRequest()

	// DELETE request
	deleteRequest();
}
