package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type trainer struct {
	Name string `bson:"Name"`
	Age  int    `bson:"Grade"`
	City string `bson:"City"`
}

// func connect() (*mongo.Database, error) {
// 	clientOptions := options.Client()
// 	clientOptions.ApplyURI("mongodb+srv://dipra123:blizzard@cluster0.euw4v.mongodb.net/dipra123?retryWrites=true&w=majority")
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = client.Connect(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client.Database("belajar_golang"), nil
// }

func connection() (*mongo.Database, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://dipra123:blizzard@cluster0.euw4v.mongodb.net/dipra123?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database("dipra123"), nil

	// clientOptions := options.Client()
	// clientOptions.ApplyURI("mongodb+srv://dipra123:blizzard@cluster0.euw4v.mongodb.net/dipra123?retryWrites=true&w=majority")
	// client, err := mongo.NewClient(clientOptions)

	// if err != nil {
	// 	return nil, err
	// }

	// err = client.Connect(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// return client.Database("belajar_golang"), nil
}

func main() {

	connection()

	port := ":8080"

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// r.Get("/trainer", func(w http.ResponseWriter, r *http.Request) {
	// 	db, err := connect()
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	_, err = db.Collection("trainer").InsertOne(ctx, trainer{"Ash", 10, "Pallet Town"})
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 		w.Write([]byte(err.Error()))
	// 	}

	// 	_, err = db.Collection("trainer").InsertOne(ctx, trainer{"Misty", 10, "Cerulen City"})
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 		w.Write([]byte(err.Error()))
	// 	}

	// 	// fmt.Println(db)
	// 	// fmt.Println(err)
	// 	w.Write([]byte("Insert success"))
	// })

	// r.Get("/trainers", func(w http.ResponseWriter, r *http.Request) {
	// 	db, err := connect()
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}

	// 	csr, err := db.Collection("trainer").Find(ctx, bson.M{"name": "Ash"})
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	defer csr.Close(ctx)

	// 	result := make([]trainer, 0)
	// 	for csr.Next(ctx) {
	// 		var row trainer
	// 		err := csr.Decode(&row)
	// 		if err != nil {
	// 			log.Fatal(err.Error())
	// 		}

	// 		result = append(result, row)
	// 	}

	// 	if len(result) > 0 {
	// 		fmt.Println("Name  :", result[0].Name)
	// 		fmt.Println("Age :", result[0].Age)
	// 	}
	// })

	// r.Get("/trainer", func(w http.ResponseWriter, r *http.Request) {
	// 	connection()
	// 	defer db.Close()
	// 	var trainers []Trainer
	// 	db.Find(&trainers)
	// 	json.NewEncoder(w).Encode(&trainers)
	// })

	// Get a handle for your collection
	// collection := client.Database("test").Collection("trainers")

	// // Some dummy data to add to the Database
	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	// // Insert a single document
	// insertResult, err := collection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// // Insert multiple documents
	// trainers := []interface{}{misty, brock}

	// insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// // Update a document
	// filter := bson.D{{"name", "Ash"}}

	// update := bson.D{
	// 	{"$inc", bson.D{
	// 		{"age", 1},
	// 	}},
	// }

	// updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// // Find a single document
	// var result Trainer

	// err = collection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Found a single document: %+v\n", result)

	// findOptions := options.Find()
	// findOptions.SetLimit(2)

	// var results []*Trainer

	// // Finding multiple documents returns a cursor
	// cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Iterate through the cursor
	// for cur.Next(context.TODO()) {
	// 	var elem Trainer
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	results = append(results, &elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// // Close the cursor once finished
	// cur.Close(context.TODO())

	// fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	// // Delete all the documents in the collection
	// deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// // Close the connection once no longer needed
	// err = client.Disconnect(context.TODO())

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("Connection to MongoDB closed.")
	// }

	fmt.Println("running in port" + port)

	http.ListenAndServe(port, r)

}
