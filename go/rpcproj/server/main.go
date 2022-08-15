package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	var (
		hm          *UserManager
		portNumber  int64 = 7000
		rpcListener net.Listener
		err         error
	)

	hm = &UserManager{}

	err = rpc.Register(hm)
	if err != nil {
		log.Fatalln(err.Error())
	}

	rpc.HandleHTTP()

	rpcListener, err = net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = http.Serve(rpcListener, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
