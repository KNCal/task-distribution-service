package setupdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"../models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "test"

// Collection Agent Assignment List
const collAgentAssignList = "assignlist"

// Collection object/instance
var collectionAgentAssignList *mongo.Collection

// Array of agents
type Agents struct {
	Agents []Agent `json:"agents"`
}

// Agent has ID and skills
type Agent struct {
	AgentID string   `json:"agentID"`
	Skills  []string `json:"skills"`
}

// Intialize table in db
func SetUpTable() {
	clientOptions := options.Client().ApplyURI(connectionString)

	ctx := context.Background()
	// connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collectionAgentAssignList = client.Database(dbName).Collection(collAgentAssignList)

	fmt.Println("collectionAgentAssignList instance created!")

	// Read agents list from file
	jsonFile, e := ioutil.ReadFile("./agents.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Println("Successfully opened agents.json")

	// map to local struct
	var agents Agents
	error := json.Unmarshal([]byte(jsonFile), &agents)
	if error != nil {
		panic(error)
	}

	// Get assignment records
	var assignment models.Assignment
	count, err1 := collectionAgentAssignList.CountDocuments(ctx, bson.M{"agentid": "1"})
	if err1 != nil {
		log.Fatal(err)
	}
	if count == 0 {
		for i := 0; i < len(agents.Agents); i++ {
			assignment.AgentID = agents.Agents[i].AgentID
			assignment.TaskID = ""
			assignment.TaskPriority = ""
			assignment.Busy = false
			assignment.TimeAssigned = time.Time{}
			insertOneAssign(assignment)
		}
	}

	err2 := client.Disconnect(ctx)
	if err2 != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
}

// Insert assignment - populate table
func insertOneAssign(assignment models.Assignment) {
	insertResult, err := collectionAgentAssignList.InsertOne(context.Background(), assignment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record into Agent Assign List", insertResult.InsertedID)
}
