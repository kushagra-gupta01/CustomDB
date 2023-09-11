package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

func New(dir string,options *Options)(*Driver, error){
	dir = filepath.Clean(dir)

	opts := Options{} 

	if opts.Logger == nil{
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir: dir,
		mutexes: make(map[string]*sync.Mutex),
		log: opts.Logger,
	}

	if _,err := os.stat(dir);err==nil{
		opts.Logger.Debug("Using %s (database already exists)\n",dir)
		return &driver,nil
	}

	opts.Logger.Debug("Creating the database at %s...\n ",dir)
	return &driver,os.MkdirAll(dir,755)
}

func (d *Driver)Write(collection,resource string,v interface{}) error{
	if collection == ""{
		return fmt.Errorf("Missing collection no place to save records")
	}

	if resource == ""{
		return fmt.Errorf("Missing resource - unable to save record(no name)!")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir,collection)
	fntPath := filepath.Join(dir,resource+"json")
	tmpPath := fntPath +".tmp"

	if err := os.MkdirAll(dir,755);err!=nil{
		return err
	}

	b,err := json.MarshalIndent(v,"","\t")
	if err!=nil{
		return err
	}
	b = append(b,byte('\n'))

	if err := os.(tmpPath,b,644);err!=nil{
		return err
	}
}

func (d *Driver)Read() error{

}

func (d *Driver)ReadAll()(){

}

func (d *Driver)Delete() error{

}

func getOrCreateMutex(collection string) *sync.Mutex{
	
}

func stat(path string)(fi os.FileInfo,err error){
	if fi,err = os.Stat(path);os.IsNotExist(err){
		fi,err = os.Stat(path+".json")
	}
	return
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

// if err :=db.Delete("user","");err !=nil{
// 	fmt.Println("Error",err)
// }