package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct{
	Title string	// title and body too needs to be exported.
	Body string		// thus Title and Body.
}

type API int

var database []Item

func (a *API) GetDB(Title string, reply *[]Item) error { // we don't really need Title string. 
	*reply = database
	return nil
}

func (a *API) GetByName(Title string, reply *Item) error{ // adding a reciever for the API pointer.
	var getItem Item
	
	for _,val := range database{
		if val.Title == Title{
			getItem=val
		}
	}
	
	*reply = getItem
	
	return nil	// but we can also return the error.
	// if there is an error then the RPC should not return any value.
}

func (a* API) AddItem(item Item, reply *Item) error{
	database = append(database, item)
	*reply = item
	return  nil
}

func (a* API) EditItem(edit Item, reply *Item) error{
	var changed Item
	
	for idx, val:=range database{
		if val.Title==edit.Title{
			database[idx]=Item{edit.Title, edit.Body}
			changed=edit
		}
	}
	
	*reply = changed
	return nil
}

func (a* API) DeleteItem(item Item, reply *Item) error{
	var del Item
	
	for idx, val := range database{
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	
	*reply = del
	return nil
}

func main(){
	fmt.Println("initial database: ", database)
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4041")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4041)
	err2 := http.Serve(listener, nil)

	if err2 != nil {
		log.Fatal("error serving: ", err)
	}
	
	/*fmt.Println("initial database: ", database)
	a := Item{"first", "a test item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}
	
	AddItem(a)
	AddItem(b)
	AddItem(c)
	
	fmt.Println("second database: ", database)
	
	DeleteItem(b)
	fmt.Println("third database: ", database)
	
	EditItem("third", Item{"fourth", "a new item"})
	fmt.Println("fourth database ", database)
	
	x:= GetByName("fourth")
	y:= GetByName("first")
	fmt.Println(x,y)*/
}
