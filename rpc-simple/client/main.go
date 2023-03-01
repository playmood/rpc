package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dial error")
	}

	var resp string
	err = client.Call("HelloService.Hello", "alice", &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	
}
