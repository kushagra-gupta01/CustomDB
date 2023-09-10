package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type(
	Logger interface{
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct{
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string
		log Logger
	}
)

type Options struct{
	Logger
}

func New()(){

}

func Write() error{

}

func Read() error{

}

func ReadAll()(){

}

func Delete() error{

}

type Address struct{
	City	string
	State string
	Country string
	Pincode json.Number
}

type User struct{
	Name string
	Age json.Number
	Contact string
	Company string
	Address Address
} 

func main(){
	dir := "./"
	db,err := New(dir,nil)
	if err!=nil{
		fmt.Println("Error",err)
	}

	employees := []User{
		{"kush","13","784568743","atlassian",Address{"lko","up","india","673990"}},
		{"kush","13","784568743","atlassian",Address{"lko","up","india","673990"}},
		{"kush","13","784568743","atlassian",Address{"lko","up","india","673990"}},
		{"kush","13","784568743","atlassian",Address{"lko","up","india","673990"}},
		{"kush","13","784568743","atlassian",Address{"lko","up","india","673990"}},
		{"kush","13","784568743","atlassian",Address{"lko","up","india","673990"}},
	}

	for _,value :=range employees{
		db.Write("users",value.Name,User{
			Name: value.Name,
			Age: value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records,err := db.ReadAll("users")
	if err != nil{
		fmt.Println("Error",err)
	}
	fmt.Println(records)

	allUsers := []User{}
	for _,f := range records{
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound);err !=nil{
			fmt.Println("Error",err)
		}
		allUsers = append(allUsers, employeeFound)
	}
	fmt.Println(allUsers)
}

// if err := db.Delete("user","kush");err !=nil{
// 	fmt.Println("Error",err)
// }

// if err :=db.DeleteALl("user","");err !=nil{
// 	fmt.Println("Error",err)
// }