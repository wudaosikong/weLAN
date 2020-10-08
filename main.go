package main

import "weLAN/server"
import "weLAN/client"

func main(){
	go client.Run()
	server.Web()
}