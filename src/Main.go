package main

import (
	"log"
	"os"
	"server"
	"time"
	"util"
)

/**
* Arguments
	* first argument is execution path
	* second argument is optional init file location

* Current implementation can be run on one or two servers.
* It's necessary to change ports in InitService to run it on two servers
*/

func main() {
	args := os.Args

	if len(args) == 1 {
		log.Println("use default init file, execution at: ", args)
		err := util.LoadFile("initFile")
		if err != nil {
			log.Fatal(err)
		}
	} else if len(args) == 2 {
		arg := os.Args[1]
		log.Println("argument for file location", arg)
		err := util.LoadFile(arg)
		if err != nil {
			log.Fatal(err)
		}
	}

	initService := server.InitService{}
	initService.Run()
	for {
		time.Sleep(10000)
	}
}
