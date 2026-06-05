package myutils

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	// "path/filepath"
	"time"
)

type Task struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetStoragePath()string{
	configDir,err := os.UserConfigDir()
	if err!=nil {
		return "task.json"
	}

	tododir := filepath.Join(configDir,"todo")
	_ = os.Mkdir(tododir,0755)

	return filepath.Join(tododir,"task.json")
}

func LoadTask()([]Task,error){
	path := GetStoragePath()
	file,err := os.ReadFile(path)
	if(err!=nil){
		fmt.Println("Error occured :",err)
		os.Exit(1)
	}
	var tasks []Task
	err = json.Unmarshal(file,&tasks)
	return tasks,err
}

func SaveTask(data []Task)([]Task,error){
	path:= GetStoragePath()
	file,err := os.Create(path)
	if err!=nil{
		fmt.Println("Error occured while loading or creating file tasks.json",err)
		os.Exit(1)
	}
	defer file.Close()
	Encoder := json.NewEncoder(file)
	Encoder.SetIndent("","  ")
	err = Encoder.Encode(data)
	if err != nil{
		fmt.Println("Error occured while writing data in tasks.json error : ",err)
		os.Exit(1)
	}
	return data,err
}