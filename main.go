package main

import (
	"github.com/fyonfa/mySocialSpace/bd" //heroku will go and find this packages
	"github.com/fyonfa/mySocialSpace/handlers"
	"log"
)


func main()  {
	if bd.CheckConnection()==0{
		log.Fatal("No Data Base connection")
		return
	}
	handlers.Handling()
}
