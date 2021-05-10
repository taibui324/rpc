package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

//calling from server
// calling http dial from previous tcp host
func main() {
	var reply Item
	var data []Item

	client, err := rpc.DialHTTP("tcp", "localhost:3000")

	if err != nil {
		log.Fatal("Error bro &d", err)
	}
	a := Item{"First", "anh tai"} // assign title by name (a, b, c)
	b := Item{"Second", "be my"}
	c := Item{"Third", "mom"}

	client.Call("API.AddItem", a, &reply) // call api from server
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &data)

	fmt.Println("Database: ", data)
	client.Call("API.EditItem", Item{"Second", "be my"}, &reply)

	client.Call("API.GetDB", "", &data)
	fmt.Println("Database: ", data)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first item: ", reply)

	client.Call("API.GetByName", "Second", &reply)
	fmt.Println("second item: ", reply)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("third item: ", reply)

}
