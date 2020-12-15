package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

)

//MongoCN, Mongo Connection is the object of connection of the BD, all the operations with BD will be use MongoCN
var MongoCN = ConnectBD()                                                                                                                             //external use, Caps, it ist going to connect to the DB and it will return the conection itself
var clientOptions = options.Client().ApplyURI("mongodb+srv://"+user+":"+pass+"@mysocial.shakx.mongodb.net/"+dataBase+"?retryWrites=true&w=majority") //we can set the URL of the DB with this

//ConnectBD is the function that allows to connect in to the DB
func ConnectBD() *mongo.Client {
	//it will connect to the database from the URL given above. It will work on the context
	//context: space memory where I will share things, executable context. like I can set mongo and this execution it cant run more than 15s. time outs
	//if the API locks/crash all the next requests will crash, then here context mas 15s for example
	//in summary, we connect here to the DB with any restrictions
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil) //ping to the DB, first we check that the DB is ok
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Successful connection!!")
	return client
}
//CheckConnection is the ping to the BD
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil) //ping to the DB, first we check that the DB is ok
	if err != nil {
		return 0
	}
	return 1
}
