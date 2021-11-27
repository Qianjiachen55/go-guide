package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Fruit struct {
	name string
	from string
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		panic(err)
	}
	collection := client.Database("mock").Collection("fruit")

	cur,err :=collection.Find(ctx,bson.D{})

	if err !=nil{
		fmt.Println("data is not exits!")
	}else {
		count := 0
		for cur.Next(ctx){
			count ++
			if count >5{
				break
			}
			var result bson.D
			err := cur.Decode(&result)
			if err != nil{
				fmt.Println("err",err)
			}else {
				fmt.Println(result)
			}
		}
	}
	peer :=bson.D{
		{"name", "peer"},
		{"from", "shanghai"},
	}
	one, err := collection.InsertOne(ctx, peer)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("res: ")
	fmt.Println(one)

}
