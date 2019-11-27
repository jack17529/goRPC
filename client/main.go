package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct{
	Title string
	Body string
}

func main(){
	var reply Item
	var db []Item
	
	client, err := rpc.DialHTTP("tcp","localhost:4041")
	
	if err != nil {
		log.Fatal("Connection Error: ", err)
	}
	
	a := Item{"first", "a first item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}
	
	client.Call("API.AddItem", a, &reply)	// we get back item in reply.
	client.Call("API.AddItem", b, &reply)	// thus we pass it as third
	client.Call("API.AddItem", c, &reply)	// parameter.
	
	// we can't access the databse locally, so we need to make a GetDB method in server
	
	client.Call("API.GetDB", "", &db)
	
	fmt.Println("Databases: ", db)
	
	client.Call("API.EditItem", Item{"second", "A new second item"}, &reply)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database: ", db)

	client.Call("API.GetByName", "first", &reply)
	fmt.Println("first item: ", reply)
}