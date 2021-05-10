package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)



type Item struct {
	Title string 
	Body string 
	
}


type API int 

var data []Item



//TODO: GETDB , Getname , edititem , delete item , update item (if possilbe)
func (a *API) GETDB(empty string, reply *[]Item) error{  //get string from database then estimate errors

    *reply = data // always recalling the main struct
    return nil

}



func (a *API) GetByName(title string, reply *Item) error { // get name to database 
	var getItem Item
	
	for _, val := range data { //data-> datbase
		if val.Title == title{ // title verification 
			getItem = val
		}
		
	}
	*reply = getItem // recalling the main struct

	return nil
}


// data = append(value)
func (a *API) AddItem(item Item, reply *Item) error { //add item to database (append data)
	data = append(data,item)
	*reply = item
	return nil
}



//syntax to remember idx, val 
// same procedure as Get name to database

func (a *API) EditItem(item Item, reply *Item) error { 

	var manipulate Item

	for idx, val := range data {
		if val.Title == item.Title { // verify real Title string
			data[idx] = Item{item.Title, item.Body} // store Title and Body
			manipulate = data[idx]
		}
	}

	*reply = manipulate
	return nil
}	


//delete item 
// going through w an algo -> either -1 or + 1 
// data = append(value)
func (a *API) DeleteItem(item Item, reply *Item) error {

	var delete Item
	

	for idx,val := range data {
		if val.Title == item.Title && val.Body == item.Body {
			data = append(data[:idx], data[idx+1:]...)

			delete = item

			break
		}
	}

	*reply = delete  // recall function as always
	return nil

}


func main () {

	//build api 
	api:= new(API) // creating api object
	err:= rpc.Register(api) // estimate error api	
	if err != nil {
		log.Fatal("can't login api", err)

	}


	//build listerner serving 
	rpc.HandleHTTP() //rpc handler http here calling port 3000
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal ("Listener error", err)
	}


	log.Println("RPC is running on port" , 3000)
	http.Serve(listener, nil)


	if err != nil {
		log.Fatal("error serving: ", err)
	}
}
